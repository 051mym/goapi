package controllers

import (
	"github.com/051mym/goapi/initializers"
	"github.com/051mym/goapi/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	// Get all records
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": "Get All Posts",
		"data":    posts,
	})
}

func GetPost(c *gin.Context) {

	// params
	id := c.Param("id")

	// Get one records
	var post models.Post
	result := initializers.DB.Find(&post, id)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": "Get Posts",
		"data":    post,
	})
}

func CreatePost(c *gin.Context) {

	var data struct {
		Title string
		Body  string
	}

	c.Bind(&data)

	post := models.Post{Title: data.Title, Body: data.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": "Create Post",
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var data struct {
		Title string
		Body  string
	}

	c.Bind(&data)

	var post models.Post
	initializers.DB.First(&post, id)

	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body:  data.Body,
	})

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": "Create Post",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": "Delete Post",
	})
}
