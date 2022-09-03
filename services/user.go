package services

import (
	"wallet/database"
	"wallet/models"
)

func AddUser(db database.IDatabase, phoneNumber string) error {
	user := models.User{
		PhoneNumber: phoneNumber,
	}

	wallet := models.Wallet{}

	if err := db.UserAdd(&user); err != nil {
		return err
	}

	wallet.UserID = user.ID
	if err := db.WalletAdd(&wallet); err != nil {
		return err
	}

	return nil
}
