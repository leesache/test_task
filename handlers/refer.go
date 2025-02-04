package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/models"
	"example.com/my_errors"

	"github.com/gin-gonic/gin"
)

// ReferUser handles applying a referral code to a user
func ReferUser(c *gin.Context) {
	// Extract user ID from the URL parameter
	id := c.Param("id")
	var userId uint
	if _, err := fmt.Sscan(id, &userId); err != nil {
		c.Error(my_errors.NewBadRequestError("Invalid user ID"))
		return
	}

	// Parse JSON input
	var input struct {
		ReferrerId uint `json:"referrer_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(my_errors.NewBadRequestError("Invalid input"))
		return
	}

	// Validate that the user exists
	_, err := models.GetUserByID(userId)
	if err != nil {
		if models.IsNotFoundError(err) {
			c.Error(my_errors.NewNotFoundError("User not found"))
		} else {
			c.Error(my_errors.NewInternalError("Failed to fetch user"))
		}
		return
	}

	// Apply referral logic
	err = models.ApplyReferral(userId, input.ReferrerId)
	if err != nil {
		if models.IsNotFoundError(err) {
			c.Error(my_errors.NewNotFoundError("User not found"))
		} else if errors.Is(err, models.ErrSelfReferral) {
			c.Error(my_errors.NewConflictError("Cannot refer yourself"))
		} else if errors.Is(err, models.ErrInvalidReferral) {
			c.Error(my_errors.NewConflictError("Invalid referral"))
		} else {
			c.Error(my_errors.NewInternalError("Failed to apply referral"))
		}
		return
	}

	// Fetch updated users
	referredUser, _ := models.GetUserByID(userId)
	referrerUser, _ := models.GetUserByID(input.ReferrerId)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message":          "Referral code applied successfully",
		"referred_balance": referredUser.Balance,
		"referrer_balance": referrerUser.Balance,
	})
}
