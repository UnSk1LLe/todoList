package service

import (
	"rest"
	"rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
	GetAll(userId int) ([]rest.TodoList, error)
	GetById(userId, listId int) (rest.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input rest.UpdateListInput) error
}
type TodoItem interface {
	Create(userId, listId int, input rest.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest.TodoItem, error)
	GetById(userId, itemId int) (rest.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rest.UpdateItemInput) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
