package auth

import "rest/models"

type UserRepository interface {
	CreateUser(*models.User) (int, error)
	GetUser(int) (*models.User, error)
	PutUser(*models.User) error
	DeleteUser(int) error
}
