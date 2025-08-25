package repo

import (
	"Store-Dio/config"
	"Store-Dio/models"
	"Store-Dio/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

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
		return false, err
	}

	return true, nil
}
func (ur *UserRepo) CheckEmailExists(email string) (bool, error) {
	stmt, err := ur.db.Prepare("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(email).Scan(&exists)

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

	err := ur.db.QueryRow(query, arg).Scan(&user.ID, &user.Email, &user.Phone, &user.Name, &user.Surname, &user.Gender, &user.Role)

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

func (ur *UserRepo) GenerateJWT(userID int, role int) (string, error) {
	exprationTime := time.Now().Add(24 * time.Hour)
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
