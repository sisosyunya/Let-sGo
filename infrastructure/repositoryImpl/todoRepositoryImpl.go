package repositoryimpl

import (
	"FirstGo/domain/repository"
	"FirstGo/graph/model"
)

type TodoRepositoryImpl struct {}

func NewTodoRepositoryImpl() repository.TodoRepository{
	return &TodoRepositoryImpl{}
}

func (repositoryImpl *TodoRepositoryImpl) FindTodo()([]*model.Todo,error) {
	// 変数指定
	db := infrastructure.GetDB()
	var err error
	// 初期化
	var todo []*model.Todo = []*model.Todo{}

	//gormを記述

	return todo,err
}


func (repositoryImpl *TodoRepositoryImpl) CreateTodo(input model.NewTodo)(*model.Todo,error){
	db := infrastructure.GetDB()
	var err error
	var todo *model.Todo = &model.Todo{}

	//gormを記述

	return todo,err
}