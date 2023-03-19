package service

import "sample/entity"

type TodoService interface {
	FindAll() []entity.Todo
	FindOne() entity.Todo
}

type todoService struct {
	todos []entity.Todo
}

func NewTodoService() TodoService {
	return &todoService{}

	// service := todoService{
	// 	todos: []entity.Todo{
	// 		{
	// 			Title:   "My First Video",
	// 			Content: "This is a video description",
	// 		},
	// 	},
	// }

	// return &service
}

func (service *todoService) FindAll() []entity.Todo {
	return service.todos
}

func (service *todoService) FindOne() entity.Todo {
	return service.todos[0]
}
