package connection

import (
	"arvan-wallet/entity"
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
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
		if err := db.AutoMigrate(&entity.User{}, &entity.UserCoupon{}, &entity.UserTxn{}); err != nil {
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

func GetPostgresInstance(config DatabaseConfig) Database {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(config.connectionString), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}
		DB = Database{
			DB: db,
		}
		if err := db.AutoMigrate(&entity.User{}, &entity.UserCoupon{}, &entity.UserTxn{}); err != nil {
			panic(errors.New("fail in migration"))
		}
	})
	return DB
}

func GetPostgresConfig() DatabaseConfig {
	return DatabaseConfig{
		connectionString: fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")),
	}
}
