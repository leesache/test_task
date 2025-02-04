package models

import "log"

// ApplyReferral applies a referral code to a user
func ApplyReferral(userId, referrerId uint) error {
	// Fetch the referred user
	referredUser, err := GetUserByID(userId)
	if err != nil {
		log.Printf("Error fetching referred user: %v", err)
		return err
	}

	// Fetch the referrer user
	referrerUser, err := GetUserByID(referrerId)
	if err != nil {
		log.Printf("Error fetching referrer user: %v", err)
		return ErrInvalidReferral
	}

	// Check for self-referral
	if userId == referrerId {
		log.Println("Error: User cannot refer themselves")
		return ErrSelfReferral
	}

	// Check if the referrer is valid (e.g., has a valid status)
	if referrerUser.Status != 0 {
		log.Println("Error: Referrer is not eligible")
		return ErrInvalidReferral
	}

	// Check if the user already has a referrer
	if referredUser.ReferrerID != 0 {
		log.Println("Error: User already has a referrer")
		return ErrInvalidReferral
	}

	// Apply the referral logic (e.g., update balances)
	referredUser.Balance += 100
	referrerUser.Balance += 50

	// Save changes to the database
	if err := DB.Save(&referredUser).Error; err != nil {
		log.Printf("Error saving referred user: %v", err)
		return err
	}
	if err := DB.Save(&referrerUser).Error; err != nil {
		log.Printf("Error saving referrer user: %v", err)
		return err
	}

	return nil
}
