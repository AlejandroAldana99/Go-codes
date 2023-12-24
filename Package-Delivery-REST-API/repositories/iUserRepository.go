package repositories

import "github.com/AlejandroAldana99/mvp_api/models"

type IUserRepository interface {
	GetUser(userID string) (models.UserData, error)
	GetUserByEmail(email string) (models.UserData, error)
	CreateUser(data models.UserData) error
}
