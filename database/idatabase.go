package database

import "wallet/models"

type IDatabase interface {
	Init()
	GetWalletById(int) (models.Wallet, error)
	UserAdd(*models.User) error
	WalletAdd(*models.Wallet) error
	WalletUpdate(*models.Wallet) error
	WalletGetByPhoneNumber(string) error
}
