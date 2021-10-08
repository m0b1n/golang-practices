package connection

import (
	"arvan-coupon/entity"
	"errors"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type DatabaseConfig struct {
	connectionString string
}

var DB Database
var once sync.Once

func GetInstance(config DatabaseConfig) Database {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(config.connectionString), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}
		DB = Database{
			DB: db,
		}
		if err := db.AutoMigrate(&entity.Coupon{}); err != nil {
			panic(errors.New("fail in migration"))
		}
	})
	return DB
}

func GetSqlLiteConfig() DatabaseConfig {
	return DatabaseConfig{
		connectionString: "gorm.db",
	}
}
