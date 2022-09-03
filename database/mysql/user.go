package mysql

import "wallet/models"

func (db *MysqlDatabase) UserAdd(user *models.User) error {
	return db.db.Save(user).Error
}
