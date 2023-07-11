package controllers

import (
	"Crud/initializers"
	"Crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	//create a post
	result := initializers.DB.Create(&post)
	// return it
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	// get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id of url
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {

	// get id from url
	id := c.Param("id")

	// get data from body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// find the post were update
	var post models.Post
	initializers.DB.First(&post, id)
	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// respond it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
