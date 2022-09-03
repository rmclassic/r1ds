package mysql

import (
	"errors"
	"wallet/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c *MysqlDatabase) GetWalletById(id int) (models.Wallet, error) {
	wallet := models.Wallet{}
	return wallet, c.db.Find(&wallet, models.Wallet{ID: id}).Error
}

func (c *MysqlDatabase) UseDiscount(walletId int, discountId int) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		wallet := models.Wallet{ID: walletId}
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}).Find(&wallet, &wallet); err != nil {
			return err.Error
		}

		discount := models.Discount{ID: discountId}
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}).Find(&discount, &discount); err != nil {
			return err.Error
		}

		if discount.QuantityLeft <= 0 {
			return errors.New("no more discounts left")
		}

		discount.QuantityLeft--
		if err := tx.Save(discount); err != nil {
			return err.Error
		}

		// TODO: Add discount usage row
		wallet.Credit += discount.Amount
		if err := tx.Save(wallet); err != nil {
			return err.Error
		}
		return nil
	})

	return err
}
