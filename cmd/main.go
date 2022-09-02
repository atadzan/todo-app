package main

import(
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/handler"
  "log"
)

func main(){
  handlers := new(handler.Handler)
  srv := new(todo.Server)
  if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
    log.Fatalf("error occured while running http server: %s", err.Error())
  }
}
