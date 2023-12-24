package services

import "github.com/AlejandroAldana99/Package-Delivery-REST-API/models"

type IUserService interface {
	GetUser(userID string) (models.UserData, error)
	CreateUser(data models.UserData) error
	Login(user string, password string) (models.LoginData, error)
}
