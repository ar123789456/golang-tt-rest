package delivery

import (
	"fmt"
	"net/http"
	"rest/substr"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase substr.UseCase
}

func NewHandler(useCase substr.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type input struct {
	Text string `json:"text"`
}

type output struct {
	Substring string `json:"substring"`
}

func (h *Handler) Post(c *gin.Context) {
	inp := new(input)
	err := c.BindJSON(inp)
	if err != nil || c.Request.Body == nil || len(inp.Text) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	fmt.Println(inp)
	c.JSON(http.StatusOK, &output{
		Substring: h.useCase.FindLongestSubstring(inp.Text),
	})
}
