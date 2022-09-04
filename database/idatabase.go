package database

import "wallet/models"

type IDatabase interface {
	Init()
	GetForUpdate(interface{}) error

	UserAdd(*models.User) error
	UserGet(*models.User) error
	UserGetByPhoneNumber(string) (*models.User, error)

	WalletAdd(*models.Wallet) error
	WalletUpdate(*models.Wallet) error
	WalletGet(*models.Wallet) error
}
