package database

import "wallet/models"

type IDatabase interface {
	Init()
	GetWalletById(int) (models.Wallet, error)
	UseDiscount(int, int) error
}
