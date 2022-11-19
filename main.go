package main

import (
	"os"

	"github.com.gezim07/go-crud/controllers"
	"github.com.gezim07/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/", controllers.PostsIndex)
	r.POST("/", controllers.PostCreate)
	r.GET("/:id", controllers.PostsGetByIndex)
	r.PUT("/:id", controllers.PostsUpdate)
	r.DELETE("/:id", controllers.PostDelete)
	r.Run(":" + port)
}
