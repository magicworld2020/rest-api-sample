package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magicworld2020/rest-api-sample/controller"
)

func main() {
	router := gin.Default()
	router.POST("/signup", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByUserID)

	// {
	// 	v1.PUT("/posts/:id", controller.UpdatePost)
	// 	v1.DELETE("/posts/:id", controller.DeletePost)
	// }

	router.Run(":8080")
}
