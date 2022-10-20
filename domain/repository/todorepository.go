package repository

import "FirstGo/graph/model"

type TodoRepository interface {
	CreateTodo(input model.NewTodo)(*model.Todo, error)
	FindTodo()([]*model.Todo, error)
}