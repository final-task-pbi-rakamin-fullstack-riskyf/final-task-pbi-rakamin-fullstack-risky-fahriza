package router

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middlewares.Auth())
	{
		authorized.PUT("/users/:userId", controllers.UpdateUser)
		authorized.DELETE("/users/:userId", controllers.DeleteUser)

		authorized.POST("/photos", controllers.CreatePhoto)
		authorized.GET("/photos", controllers.GetPhotos)
		authorized.PUT("/photos/:photoId", controllers.UpdatePhoto)
		authorized.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	return r
}