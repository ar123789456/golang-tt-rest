package auth

import "rest/models"

type UseCase interface {
	Create(string, string) (int, error)
	Get(int) (*models.User, error)
	Put(int, string, string) error
	Delete(int) error
}
