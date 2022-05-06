package database

import (
	"fmt"
	"github.com/harlesbayu/kuncie/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection(conf *config.Database) *gorm.DB {
	fmt.Printf("Try NewDatabase %s ... \n", conf.Schema)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.Username, conf.Password, conf.Name, conf.Port)
	logMode := logger.Silent
	if conf.DebugMode {
		logMode = logger.Info
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(conf.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConnections)

	return db
}
