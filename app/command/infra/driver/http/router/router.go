package router

import (
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http"
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http/handler"
)

func GetRouter(todoHandler *handler.TodoHandler) http.Router {
	var router http.Router = NewGinRouter()

	router.GET("/health_check", func(c http.Ctx) {
		c.JSON(200, `{"message": "ok"}`)
	})

	router.POST("/todos", todoHandler.Create)
	router.PUT("/todos/:uuid", todoHandler.Update)
	router.DELETE("/todos/:uuid", todoHandler.Delete)

	return router
}
