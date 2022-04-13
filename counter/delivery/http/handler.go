package handler

import (
	"net/http"
	"rest/counter"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase counter.ConterUseCase
}

func NewHandler(useCase counter.ConterUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Add(c *gin.Context) {
	istr := c.Param("i")
	i, err := strconv.Atoi(istr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = h.useCase.Add(i)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) Sub(c *gin.Context) {
	istr := c.Param("i")
	i, err := strconv.Atoi(istr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = h.useCase.Sub(i)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type counterOutput struct {
	Val int `json:"val"`
}

func (h *Handler) Val(c *gin.Context) {
	i, err := h.useCase.Val()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &counterOutput{
		Val: i,
	})
}
