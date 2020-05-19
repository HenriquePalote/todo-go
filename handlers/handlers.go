package handlers

import (
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
