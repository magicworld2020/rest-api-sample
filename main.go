package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magicworld2020/rest-api-sample/controller"
)

func main() {
	router := gin.Default()
	router.POST("/signup", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByUserID)
	router.PATCH("/users/:id", controller.UpdateUser)
	router.POST("/close", controller.DeleteUser)
	router.Run(":8080")

}
