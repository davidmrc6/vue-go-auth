package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
  api := r.Group("/api/v1")
  {
    api.POST("/login", controllers.Login)
	  api.POST("/register", controllers.Register)
	  api.GET("/home", controllers.Home)
	  api.GET("/logout", controllers.Logout)
		api.GET("/me", controllers.Me)
  }
}
