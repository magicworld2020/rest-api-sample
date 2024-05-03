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
	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully created",
		"user": gin.H{
			"user_id":  user.UserID,
			"nickname": user.Nickname,
		},
	})
}

func GetUserByUserID(c *gin.Context) {
	userID := c.Param("id")
	userService := service.UserService{}
	user, err := userService.GetUserByUserID(userID)
	if err != nil {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"user_id":  user.UserID,
			"nickname": user.Nickname,
		},
	})
}
