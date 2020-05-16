package main

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding: "required"`
}

var todoList []Todo

func main() {
	server := gin.Default()

	server.GET("/todo", func(context *gin.Context) {
		context.JSON(200, todoList)
	})

	server.POST("/todo", func(context *gin.Context) {
		var todo Todo
		context.BindJSON(&todo)
		todo.Id = 1 // TODO: implement nextId
		todoList = append(todoList, todo)

		context.Status(201)
	})

	server.Run()
}
