package main

import (
	"github.com.gezim07/go-crud/initializers"
	"github.com.gezim07/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.Db.AutoMigrate(&models.Post{})
}
