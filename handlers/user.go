package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/models"
	"example.com/my_errors"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := models.HashPassword(user.UserPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.UserPassword = hashedPassword

	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	token, err := models.GenerateJWT(user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
	})
}

// GetUsers fetches all users sorted by balance (leaderboard)
func GetUsers(c *gin.Context) {
	// Fetch all users sorted by balance
	users, err := models.GetAllUsersSortedByBalance()
	if err != nil {
		if errors.Is(err, my_errors.ErrNoUsersFound) {
			c.Error(my_errors.ErrNoUsersFound)
		} else {
			c.Error(my_errors.NewInternalError("Failed to fetch users"))
		}
		return
	}

	// Respond with the list of users
	c.JSON(http.StatusOK, gin.H{
		"message": "Leaderboard fetched successfully",
		"users":   users,
	})
}

// GetUserInfo fetches a specific user's details by ID
func GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	var userId uint
	if _, err := fmt.Sscan(id, &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := models.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
