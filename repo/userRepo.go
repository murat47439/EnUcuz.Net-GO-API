package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/golang-jwt/jwt/v5"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

// USER

func (ur *UserRepo) CreateUser(user models.User) (bool, error) {

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		config.Logger.Printf("Hash Password Error")
		return false, err
	}
	user.Password = password
	stmt, err := ur.db.Prepare("INSERT INTO USERS (email, phone, name, surname, gender, role, password) VALUES(?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Phone, user.Name, user.Surname, user.Gender, user.Role, user.Password)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ur *UserRepo) Login(email string, password string) (bool, error) {
	if email == "" || password == "" {
		return false, fmt.Errorf("Invalid data")
	}

	var hash string
	err := ur.db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hash)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("User not found")
		}
		return false, err
	}

	err = utils.CheckPasswordHash(password, hash)

	if err != nil {
		return false, fmt.Errorf("Password not correct")
	}

	return true, nil
}
func (ur *UserRepo) Logout(userID int, refreshToken string) (bool, error) {
	config.Logger.Printf("%d %s", userID, refreshToken)
	if userID == 0 || refreshToken == "" {
		return false, fmt.Errorf("Invalid data")
	}
	query := "DELETE FROM tokens WHERE token = ? AND user_id = ?"

	result, err := ur.db.Exec(query, refreshToken, userID)

	if err != nil {
		return false, fmt.Errorf("Log out unsuccessfully")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("no matching token found")
	}
	return true, nil

}
func (ur *UserRepo) Update(user models.User) (models.User, error) {
	if user.Email == "" || user.Name == "" || user.Surname == "" || user.ID == 0 {
		return models.User{}, fmt.Errorf("Invalid data")
	}
	query := "UPDATE USERS SET name=? ,surname = ? ,email = ? ,phone = ? ,gender = ? WHERE id=?"

	_, err := ur.db.Exec(query, user.Name, user.Surname, user.Email, user.Phone, user.Gender, user.ID)

	if err != nil {
		config.Logger.Printf("Failed to update user")
		return models.User{}, fmt.Errorf("Failed to update user")
	}
	return user, nil
}
func (ur *UserRepo) CheckEmailExists(email string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := ur.db.Get(&exists, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}

func (ur *UserRepo) GetUser(where string, arg any) (*models.User, error) {
	user := &models.User{}

	allowedFields := map[string]string{
		"email": "email",
		"id":    "id",
		"phone": "phone",
	}
	column, ok := allowedFields[where]
	if !ok {
		return nil, fmt.Errorf("Invalid fields : %s", where)
	}

	query := fmt.Sprintf("SELECT id, email, phone, name, surname, gender, role FROM users WHERE %s = ?", column)

	err := ur.db.Get(user, query, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) GetUserDataByEmail(email string) (*models.User, error) {
	return ur.GetUser("email", email)
}
func (ur *UserRepo) GetUserDataByID(id int) (*models.User, error) {
	return ur.GetUser("id", id)
}

// JWT
func (ur *UserRepo) NewTokens(userID int, role int) (string, string, error) {
	_, err := ur.GetUserDataByID(userID)
	if err != nil {
		return "", "", fmt.Errorf("User not found")
	}
	accessToken, err := ur.GenerateJWT(userID, role)

	if err != nil {
		return "", "", err
	}
	refreshToken, err := utils.GenerateRandomToken(32)

	if err != nil {
		return "", "", err
	}
	err = ur.StoreRefreshToken(userID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (ur *UserRepo) GenerateJWT(userID int, role int) (string, error) {
	exprationTime := time.Now().Add(15 * time.Minute)
	claims := models.JwtToken{
		UserID:   userID,
		UserRole: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exprationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(config.JWTSecret)
	if err != nil {
		config.Logger.Println("Could not create token: %v", err)
		return "", err
	}
	return tokenstring, nil
}
func (ur *UserRepo) RefreshToken(userID int, refreshToken string) (string, error) {
	var refresh models.RefreshToken

	err := ur.db.Get(&refresh, "SELECT * FROM tokens WHERE user_id = ? AND token = ?", userID, refreshToken)

	if err != nil {
		return "", fmt.Errorf("Invalid refresh token")
	}

	if time.Now().After(refresh.ExpiresAt) {
		_, err := ur.Logout(userID, refreshToken)
		if err != nil {
			config.Logger.Printf("Logout function error")
		}
		return "", fmt.Errorf("Refresh token expired")
	}

	user, err := ur.GetUserDataByID(userID)

	if err != nil {
		return "", fmt.Errorf("User not found")
	}

	newAccessToken, err := ur.GenerateJWT(user.ID, user.Role)

	if err != nil {
		return "", fmt.Errorf("Could not generate new access token")
	}

	return newAccessToken, nil
}
func (ur *UserRepo) StoreRefreshToken(userID int, refreshToken string) error {
	if userID == 0 || refreshToken == "" {
		return fmt.Errorf("Store Refresh Token Error")
	}
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES (:user_id, :token, :expires_at)"

	_, err := ur.db.NamedExec(query, map[string]interface{}{
		"user_id":    userID,
		"token":      refreshToken,
		"expires_at": expiresAt,
	})
	if err != nil {
		return err
	}
	return nil
}
func (ur *UserRepo) DecodeJWT(tokenstring models.Token) (models.JwtToken, error) {
	jwttoken := &models.JwtToken{}

	token, err := jwt.ParseWithClaims(tokenstring.Token, jwttoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return models.JwtToken{}, fmt.Errorf("User not found or Token not true")
	}
	return *jwttoken, nil
}

//ADMİN CONTROL

func (ur *UserRepo) OnlyAdmin(userID int) (bool, error) {
	stmt, err := ur.db.Prepare("SELECT 1 FROM users WHERE role=1 AND id = ?")

	if err != nil {

		return false, err
	}
	defer stmt.Close()

	var tmp int

	err = stmt.QueryRow(userID).Scan(&tmp)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil

}
