package models

func CompleteTask(userID uint) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := GetUserByID(userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	user.Balance += 100
	if err := UpdateUser(user); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
