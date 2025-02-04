package handlers

import (
	"fmt"
	"net/http"

	"example.com/models"
	"example.com/my_errors"

	"github.com/gin-gonic/gin"
)

// TaskComplete handles completing a task and updating the user's balance
func TaskComplete(c *gin.Context) {
	id := c.Param("id")
	var userId uint
	if _, err := fmt.Sscan(id, &userId); err != nil {
		c.Error(my_errors.NewBadRequestError("Invalid user ID"))
		return
	}

	err := models.CompleteTask(userId)
	if err != nil {
		if models.IsNotFoundError(err) {
			c.Error(my_errors.NewNotFoundError("User not found"))
		} else if models.IsConflictError(err) {
			c.Error(my_errors.NewConflictError("Task already completed"))
		} else {
			c.Error(my_errors.NewInternalError("Failed to complete task"))
		}
		return
	}

	user, _ := models.GetUserByID(userId)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Task completed successfully",
		"user_balance": user.Balance,
	})
}
