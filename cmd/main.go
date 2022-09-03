package main

import(
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/handler"
  "github.com/atadzan/todo-app/pkg/repository"
  "github.com/atadzan/todo-app/pkg/service"
  "log"
)

func main(){
  repos := repository.NewRepository()
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  srv := new(todo.Server)
  if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
    log.Fatalf("error occured while running http server: %s", err.Error())
  }
}
