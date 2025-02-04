package models

import "example.com/my_errors"

type User struct {
	UserId       uint   `gorm:"primaryKey" json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	Balance      uint   `json:"balance"`
	Status       int8   `json:"status"`
	ReferrerID   uint   `json:"referrer_id"`
}

// CreateUser saves a new user to the database
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uint) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}

// UpdateUser updates a user's details in the database
func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.Where("user_name = ?", username).First(&user).Error
	return &user, err
}

// GetAllUsers retrieves all users sorted by balance
func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Order("balance DESC").Find(&users).Error
	return users, err
}

// GetAllUsersSortedByBalance fetches all users sorted by balance in descending order
func GetAllUsersSortedByBalance() ([]User, error) {
	var users []User
	result := DB.Order("balance DESC").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return nil, my_errors.ErrNoUsersFound
	}
	return users, nil
}
