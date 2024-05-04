package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
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

func authenticateUser(authHeader string) bool {
	if authHeader == "" {
		return false
	}

	// Extract the credentials from the Authorization header
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Basic" {
		return false
	}

	encodedCredentials := parts[1]

	// Decode the Base64-encoded credentials
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return false
	}

	// Split the decoded credentials into user_id and password
	credentials := strings.SplitN(string(decodedCredentials), ":", 2)
	if len(credentials) != 2 {
		fmt.Println("Invalid credentials format")
		return false
	}

	userID := credentials[0]
	password := credentials[1]
	userService := service.UserService{}
	return authenticateUserWithCredentials(userID, password, &userService)
}

func authenticateUserWithCredentials(userID, password string, userService *service.UserService) bool {
	user, err := userService.GetUserByUserID(userID)
	if err != nil {
		// User not found or other error occurred
		return false
	}

	// Check if the provided password matches the user's password
	return user != nil && user.Password == password
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

func GetUserByUserID(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	// Authenticate user
	if !authenticateUser(authHeader) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
		return
	}

	userID := c.Param("id")
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

func UpdateUser(c *gin.Context) {
	// Parse Authorization header
	authHeader := c.GetHeader("Authorization")

	// Authenticate user
	if !authenticateUser(authHeader) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
		return
	}

	userID := c.Param("id")

	// Check if the user exists
	userService := service.UserService{}
	user, err := userService.GetUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User found"})
		return
	}

	// Bind request body to struct representing update payload
	var updateUser model.User
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "cause": err.Error()})
		return
	}

	// Check if either nickname or comment is provided
	if updateUser.Nickname == "" && updateUser.Comment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User updation failed", "cause": "required nickname or comment"})
		return
	}

	// Check if user_id or password is being changed
	if updateUser.UserID != "" || updateUser.Password != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not change user_id and password", "cause": "not updatable user_id and password"})
		return
	}

	// Update user information if provided
	if updateUser.Nickname != "" {
		user.Nickname = updateUser.Nickname
	} else {
		user.Nickname = user.UserID
	}
	if updateUser.Comment != "" {
		user.Comment = updateUser.Comment
	} else {
		user.Comment = ""
	}
	// Specifying a user with an ID different from the authenticated user
	parts := strings.Split(authHeader, " ")
	encodedCredentials := parts[1]
	// Decode the Base64-encoded credentials
	decodedCredentials, _ := base64.StdEncoding.DecodeString(encodedCredentials)
	// Split the decoded credentials into user_id and password
	credentials := strings.SplitN(string(decodedCredentials), ":", 2)
	authUserID := credentials[0]
	if authUserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "No Permission for Update"})
		return
	}

	// Save updated user information to the database
	err = userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully updated",
		"user": gin.H{
			"nickname": user.Nickname,
			"comment":  user.Comment,
		},
	})
}
