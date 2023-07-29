package data

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupDatabase(verboseLogging bool) error {
	config := &gorm.Config{}
	if verboseLogging {
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Millisecond * 200,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		)
		config = &gorm.Config{
			Logger: gormLogger,
		}
	}

	database, err := gorm.Open(sqlite.Open("data/database.db"), config)
	if err != nil {
		return err
	}

	DB = database
	return nil
}

func MigrateDatabase() error {
	err := DB.AutoMigrate(&Time{}, &Client{}, &Project{})
	if err != nil {
		return err
	}

	return nil
}
