package localstorage

import (
	"errors"
	"rest/models"
	"sync"
)

type UserLocalRepository struct {
	CurrentID int
	DB        map[int]*models.User
	Mutex     *sync.Mutex
}

func NewUserRepository(db map[int]*models.User, m *sync.Mutex) *UserLocalRepository {
	return &UserLocalRepository{
		CurrentID: 1,
		DB:        db,
		Mutex:     m,
	}
}

func (ur *UserLocalRepository) CreateUser(user *models.User) (int, error) {
	ur.Mutex.Lock()
	user.ID = ur.CurrentID
	ur.DB[ur.CurrentID] = user
	ur.CurrentID++
	ur.Mutex.Unlock()
	return user.ID, nil
}
func (ur *UserLocalRepository) GetUser(id int) (*models.User, error) {
	ur.Mutex.Lock()
	curr := ur.DB[id]
	ur.Mutex.Unlock()
	if curr == nil {
		return curr, errors.New("user not found")
	}
	return curr, nil
}
func (ur *UserLocalRepository) PutUser(user *models.User) error {
	var err error
	ur.Mutex.Lock()
	if ur.DB[user.ID] == nil {
		err = errors.New("user not found")
	} else {
		ur.DB[user.ID] = user
	}
	ur.Mutex.Unlock()
	return err
}
func (ur *UserLocalRepository) DeleteUser(id int) error {
	var err error
	ur.Mutex.Lock()
	if ur.DB[id] == nil {
		err = errors.New("user not found")
	} else {
		delete(ur.DB, id)
	}
	ur.Mutex.Unlock()
	return err
}
