package handler

import "github.com/gin-gonic/gin"

type Handler struct{

}

func (h *Handler) InitRoutes() *gin.Engine{
  router := gin.New()

  auth := router.Group(relativePath: "/auth")
  {
    auth.POST(relativePath: "/sign-up", h.signUp)
    auth.POST(relativePath: "/sign-in", h.signIn)
  }

  api := router.Group(relativePath: "/api")
  {
    lists := api.Group(relativePath:"/lists")
    {
      lists.POST(relativePath: "/", h.createList)
      lists.GET(relativePath: "/", h.getAllLists)
      lists.GET(relativePath: "/:id", h.getListById)
      lists.PUT(relativePath: "/:id", h.updateList)
      lists.DELETE(relativePath: "/:id", h.deleteList)

      items := lists.Group(relativePath:":id/items")
      {
        items.POST(relativePath: "/", h.createItem)
        items.GET(relativePath: "/", h.getAllItems)
        items.GET(relativePath: "/:items_id", h.getItemById)
        items.PUT(relativePath: "/:items_id", h.updateItem)
        items.DELETE(relativePath: "/:items_id", h.deleteItem)
      }
    }
  }
  return router
}
