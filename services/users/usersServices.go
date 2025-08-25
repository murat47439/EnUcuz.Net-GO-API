package users

import (
	"Store-Dio/models"
	"Store-Dio/repo"
	"fmt"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {

	if user.Name == "" || user.Email == "" || user.Surname == "" || user.Password == "" {
		return models.User{}, fmt.Errorf("Some data is empty")
	}

	existEmail, err := s.UserRepo.CheckEmailExists(user.Email)

	if err != nil {
		return models.User{}, fmt.Errorf("CheckEmailExists error : %v", err)
	}
	if existEmail {
		return models.User{}, fmt.Errorf("Email already exists")
	}

	_, err = s.UserRepo.CreateUser(user)

	if err != nil {
		return models.User{}, fmt.Errorf("Create User error: %v", err)
	}

	return user, nil
}
func (s *UserService) Login(user models.User) (string, error) {

	_, err := s.UserRepo.Login(user.Email, user.Password)

	if err != nil {
		return "", err
	}
	userdata := &models.User{}
	userdata, err = s.UserRepo.GetUserDataByEmail(user.Email)

	if err != nil {
		return "", err
	}

	jwttoken, err := s.UserRepo.GenerateJWT(userdata.ID, userdata.Role)

	if err != nil {
		return "", err
	}

	return jwttoken, nil

}
