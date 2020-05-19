package handlers

import (
	"strconv"

	"github.com/HenriquePalote/todo-go/domains"
	"github.com/HenriquePalote/todo-go/utils"
	"github.com/gin-gonic/gin"
)

var todoList []domains.Todo

func GetTodoList(context *gin.Context) {
	context.JSON(200, todoList)
}

func CreateTodo(context *gin.Context) {
	var todo domains.Todo
	context.BindJSON(&todo)

	todo.Id = utils.GetNextId(todoList)
	todoList = append(todoList, todo)

	context.JSON(201, todo.Id)
}

func GetTodo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	for _, todo := range todoList {
		if todo.Id == id {
			context.JSON(200, todo)
			return
		}
	}

	context.JSON(404, nil)
}
