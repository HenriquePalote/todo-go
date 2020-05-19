package server

import (
	"github.com/HenriquePalote/todo-go/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) *gin.Engine {
	
	server.GET("/todo", handlers.GetTodoList)

	server.POST("/todo", handlers.CreateTodo)

	return server
}