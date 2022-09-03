package services

import (
	"wallet/database"
	"wallet/models"
)

func GetWalletBalance(db database.IDatabase, walletId int) (models.Wallet, error) {
	return db.GetWalletById(walletId)
}

func ChargeWallet(db database.IDatabase, phoneNumber string, amount float64) error {
	db.WalletGetByPhoneNumber(phoneNumber)
	db.WalletUpdate()

}
