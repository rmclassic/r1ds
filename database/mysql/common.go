package mysql

import "gorm.io/gorm/clause"

func (db *MysqlDatabase) GetForUpdate(item interface{}) error {
	return db.db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Find(item, item).Error
}
