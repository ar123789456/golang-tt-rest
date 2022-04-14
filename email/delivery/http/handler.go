package handler

import (
	"io/ioutil"
	"net/http"
	"rest/email"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	useCase email.EmailUseCase
}

func NewEmailHandler(useCase email.EmailUseCase) *EmailHandler {
	return &EmailHandler{
		useCase: useCase,
	}
}

func (eh *EmailHandler) CheckEmail(c *gin.Context) {
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := eh.useCase.ValidEmail(string(jsonDataBytes))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, res)
}
