package app

import (
	"fmt"
	"os"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/helper"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.Err(err)

	fileLogger, err := helper.NewFileGormLogger(logger.Info, "pkg/log/gorm.log")
	helper.Err(err)

	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger:                 fileLogger,
		DisableAutomaticPing:   true,
		SkipDefaultTransaction: true,
	})
	helper.Err(err)

	// if os.Getenv("MODE") == "development" {
	// 	resetDatabase(db, &model.User{})
	// 	seed.MigrateUserData(10, db)
	// } else if os.Getenv("MODE") == "release" {
	// 	db.AutoMigrate(&model.User{})
	// }

	db.AutoMigrate(
		&model.User{},
		&model.TokenAuth{},
		&model.Group{},
		&model.SmartBin{},
		&model.Config{},
		&model.History{},
		&model.Notification{},
	)

	return db
}
