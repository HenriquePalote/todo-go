package utils

import "github.com/HenriquePalote/todo-go/domains"

func GetNextId(list []domains.Todo) int {
	arrayLen := len(list)

	if arrayLen > 0 {
		nextId := list[arrayLen-1].Id + 1
		return nextId
	}

	return 1
}
