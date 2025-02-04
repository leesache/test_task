package handlers

import (
	"net/http"

	"example.com/models"
	"example.com/my_errors"
	"github.com/gin-gonic/gin"
)

// LoginUser handles user login and generates a JWT token
func LoginUser(c *gin.Context) {
	var input struct {
		UserName string `json:"user_name"`
		Password string `json:"user_password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(my_errors.NewBadRequestError("Invalid input"))
		return
	}

	user, err := models.GetUserByUsername(input.UserName)
	if err != nil {
		if models.IsNotFoundError(err) { // Check if the error is "user not found"
			c.Error(my_errors.NewUnauthorizedError("Invalid credentials"))
		} else {
			c.Error(my_errors.NewInternalError("Failed to fetch user"))
		}
		return
	}

	if err := models.VerifyPassword(user.UserPassword, input.Password); err != nil {
		c.Error(my_errors.NewUnauthorizedError("Invalid credentials"))
		return
	}

	token, err := models.GenerateJWT(user.UserId)
	if err != nil {
		c.Error(my_errors.NewInternalError("Error generating token"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
