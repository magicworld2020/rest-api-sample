package controller

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/magicworld2020/rest-api-sample/model"
	"github.com/magicworld2020/rest-api-sample/service"

	"github.com/gin-gonic/gin"
)

func isAlphanumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
func CreateUser(c *gin.Context) {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Account creation failed", "cause": "required user_id and password"})
		return
	}

	// Check if user_id or password is empty
	if user.UserID == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Account creation failed", "cause": "user_id and password are required"})
		return
	}

	// Check user_id and password length
	if len(user.UserID) < 6 || len(user.UserID) > 20 || len(user.Password) < 8 || len(user.Password) > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Account creation failed", "cause": "invalid user_id or password length"})
		return
	}

	// Check user_id and password pattern (alphanumeric characters only)
	if !isAlphanumeric(user.UserID) || !isAlphanumeric(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Account creation failed", "cause": "invalid user_id or password format"})
		return
	}

	userService := service.UserService{}

	// Check if user_id already exists
	existingUser, err := userService.GetUserByUserID(user.UserID)
	if existingUser != nil || err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Account creation failed", "cause": "already same user_id is used"})
		return
	}

	// Create new user
	if err := userService.AddUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account successfully created", "user": gin.H{"user_id": user.UserID, "nickname": user.Nickname}})
}

// func GetUserByUserID(c *gin.Context) {
// 	userID := c.Param("id")
// 	userService := service.UserService{}
// 	user, err := userService.GetUserByUserID(userID)
// 	if err != nil {
// 		c.String(http.StatusNotFound, "User not found")
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"user": gin.H{
// 			"user_id":  user.UserID,
// 			"nickname": user.Nickname,
// 		},
// 	})
// }

func GetUserByUserID(c *gin.Context) {
	userID := c.Param("id")

	// Authenticate user
	if !authenticateUser(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
		return
	}

	userService := service.UserService{}
	user, err := userService.GetUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User found"})
		return
	}
	fmt.Println(user.Nickname)
	if user.Nickname != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "User details by user_id",
			"user": gin.H{
				"user_id":  user.UserID,
				"nickname": user.Nickname,
				"comment":  user.Comment,
			},
		})
	} else {
		nickname := user.UserID // Set nickname as user_id if it's not set
		c.JSON(http.StatusOK, gin.H{
			"message": "User details by user_id",
			"user": gin.H{
				"user_id":  user.UserID,
				"nickname": nickname,
			},
		})
	}
}

func authenticateUser(c *gin.Context) bool {
	// Implement authentication logic here, return true if authenticated, false otherwise
	return true // Placeholder, replace with actual authentication logic
}
