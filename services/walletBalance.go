package services

import (
	"wallet/database"
	"wallet/models"
)

func GetWalletBalance(db database.IDatabase, walletId int) (models.Wallet, error) {
	return db.GetWalletById(walletId)
}

func UseDiscount(db database.IDatabase, walletId int, discountId int) error {
	return db.UseDiscount(walletId, discountId)
}
