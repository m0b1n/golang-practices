package sqlLiteRepo

import (
	"arvan-wallet/dal/connection"
	"arvan-wallet/entity"
	"gorm.io/gorm"
	"log"
)

type UserTxnRepository struct {
	DB connection.Database
}

func (u UserTxnRepository) SaveUserTxn(txn entity.UserTxn) error {
	result := u.DB.Create(&txn)
	return result.Error
}

func (u UserTxnRepository) GetUserTxnsByUser(user entity.User) ([]entity.UserTxn, error) {
	var userTxns []entity.UserTxn
	result := u.DB.Model(&entity.UserTxn{}).Joins("join users on user_txns.user_id = users.id and users.phone_number = ?", user.PhoneNumber).Scan(&userTxns)
	return userTxns, result.Error
}

func NewUserTxnRepository() UserTxnRepository {
	return UserTxnRepository{
		DB: connection.GetInstance(connection.GetSqlLiteConfig()),
	}
}

func NewTxnUserTxnRepository(db *gorm.DB) UserTxnRepository {
	if db == nil {
		log.Print("Transaction Database not found")
		return UserTxnRepository{}
	}
	return UserTxnRepository{
		DB: connection.Database{
			DB: db,
		},
	}
}
