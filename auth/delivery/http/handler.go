package http

import (
	"net/http"
	"rest/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type CreatOrPutInput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type CreatOutput struct {
	ID int `json:"id"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(CreatOrPutInput)
	err := c.BindJSON(inp)

	if err != nil || len(inp.First_name) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := h.useCase.Create(inp.First_name, inp.Last_name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &CreatOutput{
		ID: id,
	})
}

type GetOutput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

func (h *Handler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := h.useCase.Get(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &GetOutput{
		First_name: user.FirstName,
		Last_name:  user.LastName,
	})
}

func (h *Handler) Put(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	inp := new(CreatOrPutInput)
	if err = c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = h.useCase.Put(id, inp.First_name, inp.Last_name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = h.useCase.Delete(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
