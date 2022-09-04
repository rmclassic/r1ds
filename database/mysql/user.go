package mysql

import "wallet/models"

func (db *MysqlDatabase) UserAdd(user *models.User) error {
	return db.db.Save(user).Error
}

func (db *MysqlDatabase) UserGetByPhoneNumber(phoneNumber string) (*models.User, error) {
	user := &models.User{
		PhoneNumber: phoneNumber,
	}

	return user, db.db.Find(user, user).Error
}

func (db *MysqlDatabase) UserGet(user *models.User) error {
	return db.db.Find(user, user).Error
}
