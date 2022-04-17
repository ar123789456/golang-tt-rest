package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"rest/auth/repository/localstorage"
	"rest/auth/usecase"
	"rest/models"
	"strconv"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPost(t *testing.T) {
	i := CreatOrPutInput{
		First_name: "hello",
		Last_name:  "world",
	}

	creattestCase := map[int]interface{}{
		http.StatusOK:         i,
		http.StatusBadRequest: nil,
	}

	r := gin.Default()
	usecase := usecase.NewAuthUseCase(localstorage.NewUserRepository(map[int]*models.User{}, &sync.Mutex{}))

	RegisterUserEndpoints(r, usecase)

	for code, inp := range creattestCase {
		body, err := json.Marshal(inp)
		if err != nil {
			fmt.Println("jmo:", err)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/rest/user", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)
		defer req.Body.Close()
		_, err = io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
		assert.Equal(t, code, w.Code)
	}

}

func TestGet(t *testing.T) {
	i := CreatOrPutInput{
		First_name: "hello",
		Last_name:  "world",
	}

	r := gin.Default()
	usecase := usecase.NewAuthUseCase(localstorage.NewUserRepository(map[int]*models.User{}, &sync.Mutex{}))

	RegisterUserEndpoints(r, usecase)
	id, _ := usecase.Create(i.First_name, i.Last_name)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/rest/user/"+strconv.Itoa(id), nil)
	r.ServeHTTP(w, req)
	defer req.Body.Close()

	assert.Equal(t, http.StatusOK, w.Code)
}
