package services

import (
	"wallet/database"
	"wallet/models"
)

func AddUser(db database.IDatabase, user *models.User) error {
	wallet := models.Wallet{}

	if err := db.UserAdd(user); err != nil {
		return err
	}

	wallet.UserID = user.ID
	if err := db.WalletAdd(&wallet); err != nil {
		return err
	}

	return nil
}

func GetUserByPhoneNumber(db database.IDatabase, phoneNumber string) (*models.User, error) {
	user := &models.User{
		PhoneNumber: phoneNumber,
	}

	return user, db.UserGet(user)
}
