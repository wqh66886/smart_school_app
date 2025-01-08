package db

import (
	"fmt"
	"github.com/wqh/smart/school/system/internal/domain"
	"log"

	"github.com/wqh/smart/school/system/internal/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(db *gorm.DB, config configs.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database)
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	}
	db.AutoMigrate(domain.User{}, domain.School{})
}
