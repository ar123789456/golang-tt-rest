package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"rest/substr/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPost(t *testing.T) {
	i := &input{
		Text: "re",
	}
	testCase := map[int]interface{}{
		http.StatusOK:         i,
		http.StatusBadRequest: nil,
	}

	r := gin.Default()

	usecase.NewSubstrUseCase()

	usc := usecase.NewSubstrUseCase()

	r.POST("/rest/substr/find", NewHandler(usc).Post)

	for code, inp := range testCase {
		body, err := json.Marshal(inp)
		if err != nil {
			fmt.Println("jmo:", err)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/rest/substr/find", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)
		defer req.Body.Close()
		b, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
		assert.Equal(t, code, w.Code)
	}

}
