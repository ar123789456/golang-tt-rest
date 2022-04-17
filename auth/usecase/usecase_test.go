package usecase

import (
	"rest/auth/repository/localstorage"
	"rest/models"
	"sync"
	"testing"
)

func TestUser(t *testing.T) {
	usecase := NewAuthUseCase(localstorage.NewUserRepository(map[int]*models.User{}, &sync.Mutex{}))
	//Create test
	first_name := "Hello"
	last_name := "World"
	id, err := usecase.Create(first_name, last_name)
	if err != nil {
		t.Error(err)
	}
	if id != 1 {
		t.Errorf("Expected '%v', but got '%v'", 1, id)
	}
	//Get TEst
	user, err := usecase.Get(id)
	if err != nil {
		t.Error(err)
	}
	if user.FirstName != first_name || user.LastName != last_name {
		t.Errorf("Expected '%v %v', but got '%v %v'", first_name, last_name, user.FirstName, user.LastName)
	}
	_, err = usecase.Get(100)
	if err == nil {
		t.Errorf("Expected '%v', but got '%v'", "user not found", err)
	}
	//Put Test
	put_lastName := "w"
	usecase.Put(id, first_name, put_lastName)
	user, err = usecase.Get(id)
	if err != nil {
		t.Error(err)
	}
	if user.FirstName != first_name || user.LastName != put_lastName {
		t.Errorf("Expected '%v %v', but got '%v %v'", first_name, put_lastName, user.FirstName, user.LastName)
	}
	//delet Test
	err = usecase.Delete(id)
	if err != nil {
		t.Error(err)
	}
	_, err = usecase.Get(100)
	if err == nil {
		t.Errorf("Expected '%v', but got '%v'", "user not found", err)
	}
}
