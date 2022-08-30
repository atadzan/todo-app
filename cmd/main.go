package main

import(
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/handler"
  "log"
)

func main(){
  handlers := new(handler.Handler)
  srv := new(todo.Server)
  if err := srv.Run(port: "8000", handlers.InitRoutes()); err != nil {
    log.Fatalf(format: "error occured while running http server: %s", err.Error())
  }
}
