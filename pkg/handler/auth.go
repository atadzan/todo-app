package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/atadzan/todo-app"
  )

func (h *Handler) signUp(c *gin.Context) {
  var input todo.User

  if err := c.BindJSON(&input); err != nil {
    
  }
}

func (h *Handler) signIn(c *gin.Context){

}
