package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magicworld2020/rest-api-sample/controller"
)

func main() {
	// Load TLS certificate and private key
	// certFile := "demo/certificate.crt"
	// keyFile := "demo/private.key"
	router := gin.Default()
	router.POST("/signup", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByUserID)
	router.PATCH("/users/:id", controller.UpdateUser)
	router.POST("/close", controller.DeleteUser)
	router.Run(":8080")
	// Run HTTPS server
	// err := router.RunTLS(":443", certFile, keyFile)
	// if err != nil {
	// 	log.Fatalf("Failed to start HTTPS server: %v", err)
	// }
}
