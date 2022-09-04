package services

import (
	"wallet/database"
	"wallet/models"
)

func ChargeWallet(db database.IDatabase, userId int, amount float64) error {
	user := &models.User{
		ID: userId,
	}

	if err := db.UserGet(user); err != nil {
		return err
	}

	wallet := &models.Wallet{
		UserID: user.ID,
	}

	if err := db.GetForUpdate(wallet); err != nil {
		return err
	}

	wallet.Balance += amount

	if err := db.WalletUpdate(wallet); err != nil {
		return err
	}

	return nil
}

func GetUserWallet(db database.IDatabase, userId int) (*models.Wallet, error) {
	wallet := &models.Wallet{
		UserID: userId,
	}

	return wallet, db.WalletGet(wallet)
}
