package service

import (
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/repository"
)

type TodoListService struct{
  repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
  return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
  return s.repo.Create(userId, list)
}
