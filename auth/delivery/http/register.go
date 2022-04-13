package handler

import (
	"rest/auth"

	"github.com/gin-gonic/gin"
)

func RegisterUserEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	userEndpoints := router.Group(" /rest/user")
	{
		userEndpoints.POST("", h.Get)
		userEndpoints.GET(":id", h.Get)
		userEndpoints.PUT(":id", h.Put)
		userEndpoints.DELETE(":id", h.Delete)
	}
}
