package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magicworld2020/rest-api-sample/controller"
)

func main() {
	router := gin.Default()
	router.POST("/signup", controller.CreateUser)
	router.GET("/users/:id", controller.GetUser)

	// v1 := router.Group("/api/v1")
	// {
	// 	v1.GET("/users", controller.GetUsers)
	// 	v1.GET("/users/:id", controller.GetUser)
	// 	v1.POST("/users", controller.CreateUser)
	// 	v1.PUT("/users/:id", controller.UpdateUser)
	// 	v1.DELETE("/users/:id", controller.DeleteUser)

	// 	v1.GET("/posts", controller.GetPosts)
	// 	v1.GET("/posts/:id", controller.GetPost)
	// 	v1.POST("/posts", controller.CreatePost)
	// 	v1.PUT("/posts/:id", controller.UpdatePost)
	// 	v1.DELETE("/posts/:id", controller.DeletePost)
	// }

	router.Run(":8080")
}
