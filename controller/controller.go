package controller

import (
	"net/http"

	"github.com/magicworld2020/rest-api-sample/model"
	"github.com/magicworld2020/rest-api-sample/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.AddUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func UpdateUser(c *gin.Context) {
	// Implement logic to update an existing user
}

func DeleteUser(c *gin.Context) {
	// Implement logic to delete a user
}

func GetPosts(c *gin.Context) {
	// Implement logic to fetch posts
}

func GetPost(c *gin.Context) {
	// Implement logic to fetch a single post
}

func CreatePost(c *gin.Context) {
	// Implement logic to create a new post
}

func UpdatePost(c *gin.Context) {
	// Implement logic to update an existing post
}

func DeletePost(c *gin.Context) {
	// Implement logic to delete a post
}
