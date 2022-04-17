package http

import (
	"rest/counter"

	"github.com/gin-gonic/gin"
)

func RegisterCounterEndpoints(router *gin.Engine, uc counter.ConterUseCase) {
	h := NewHandler(uc)

	counterEndpoints := router.Group("/rest/counter")
	{
		counterEndpoints.POST("/add/:i", h.Add)
		counterEndpoints.POST("/sub/:i", h.Sub)
		counterEndpoints.GET("/val", h.Val)
	}
}
