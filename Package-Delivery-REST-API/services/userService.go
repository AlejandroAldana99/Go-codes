package services

import (
	e "errors"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository repositories.IUserRepository
}

func (service UserService) GetUser(userID string) (models.UserData, error) {
	user, err := service.Repository.GetUser(userID)
	if err != nil {
		logger.Error("services", "GetUser", err.Error())
		return user, errors.HandleServiceError(err)
	}

	return user, nil
}

func (service UserService) CreateUser(data models.UserData) error {
	// Encryp password
	pass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("services", "Createuser", err.Error())
		return errors.HandleServiceError(err)
	}

	data.Password = string(pass)
	err = service.Repository.CreateUser(data)
	if err != nil {
		logger.Error("services", "Createuser", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}

func (service UserService) Login(user string, password string) (models.LoginData, error) {
	var login models.LoginData
	userData, err := service.Repository.GetUser(user)
	if err != nil {
		userData, err = service.Repository.GetUserByEmail(user)
		if err != nil {
			logger.Error("services", "GetUser", err.Error())
			return login, errors.HandleServiceError(err)
		}
	}

	cErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if cErr != nil && cErr == bcrypt.ErrMismatchedHashAndPassword {
		err := e.New("Invalid login credentials")
		logger.Error("services", "Login", err.Error())
		return login, errors.HandleServiceError(err)
	}

	tokenString := libs.EncodeJWT(string(userData.UserID), userData.Role)

	login.Token = tokenString
	login.Status = constants.GeneralStatus

	return login, nil
}
