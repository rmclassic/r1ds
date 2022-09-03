package mysql

import (
	"os"
	"time"
	"wallet/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDatabase struct {
	db *gorm.DB
}

func (dbInstance *MysqlDatabase) Init() {
	for {
		_db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_CONN")), &gorm.Config{})
		if err != nil {
			println("failed to connect database, retrying after 10s")
			time.Sleep(time.Second * 10)
			continue
		}

		dbInstance.db = _db
		dbInstance.db.AutoMigrate(models.User{})
		dbInstance.db.AutoMigrate(models.Wallet{})
		dbInstance.db.AutoMigrate(models.Discount{})
		break
	}

}
