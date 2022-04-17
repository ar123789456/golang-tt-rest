package http

import (
	"rest/email"

	"github.com/gin-gonic/gin"
)

func RegisterEmailEndpoints(router *gin.Engine, uc email.EmailUseCase) {
	h := NewEmailHandler(uc)

	router.POST("/rest/email", h.CheckEmail)
}
