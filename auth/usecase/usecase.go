package usecase

import (
	"rest/auth"
	"rest/models"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

func NewAuthUseCase(repo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: repo,
	}
}

func (a *AuthUseCase) Create(first_name, last_name string) (int, error) {
	user := models.User{
		FirstName: first_name,
		LastName:  last_name,
	}
	return a.userRepo.CreateUser(&user)
}

func (a *AuthUseCase) Get(id int) (*models.User, error) {
	return a.userRepo.GetUser(id)
}

func (a *AuthUseCase) Put(id int, first_name, last_name string) error {
	user := models.User{
		ID:        id,
		FirstName: first_name,
		LastName:  last_name,
	}
	return a.userRepo.PutUser(&user)
}

func (a *AuthUseCase) Delete(id int) error {
	return a.userRepo.DeleteUser(id)
}
