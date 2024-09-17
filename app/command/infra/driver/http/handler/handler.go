package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lazaro-Barros/boilerplate-golang/command/application"
	http_app "github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http"
)

type TodoHandler struct {
	TodoService *application.TodoService
}

func NewTodoHandler(todoService *application.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

func (h *TodoHandler) Create(c http_app.Ctx) {
	input := application.TodoIn{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.ContextRequest()
	uuid, err := h.TodoService.Create(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": uuid})
}

func (h *TodoHandler) Update(c http_app.Ctx) {
	uuid := c.Param("uuid")
	input := application.TodoIn{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.ContextRequest()
	if err := h.TodoService.Update(ctx, uuid, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *TodoHandler) Delete(c http_app.Ctx) {
	uuid := c.Param("uuid")
	ctx := c.ContextRequest()
	if err := h.TodoService.Delete(ctx, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
