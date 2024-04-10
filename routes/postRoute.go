package routes

import (
	"github.com/051mym/goapi/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}
