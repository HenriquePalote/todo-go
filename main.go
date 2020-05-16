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

		todo.Id = getNextId(todoList)
		todoList = append(todoList, todo)

		context.JSON(201, todo.Id)
	})

	server.Run()
}

func getNextId(list []Todo) int {
	arrayLen := len(todoList)

	if arrayLen > 0 {
		nextId := list[arrayLen-1].Id + 1
		return nextId
	}

	return 1
}
