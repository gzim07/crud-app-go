package controllers

import (
	"github.com.gezim07/go-crud/initializers"
	"github.com.gezim07/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.Db.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	})
}
func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.Db.Find(&posts)
	c.JSON(200, gin.H{
		"post": posts,
	})
}
func PostsGetByIndex(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.Db.First(&post, "ID = ?", id)
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	var post models.Post
	initializers.Db.First(&post, "ID = ?", id)
	initializers.Db.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.Db.Delete(&models.Post{}, id)
	c.Status(200)
}
