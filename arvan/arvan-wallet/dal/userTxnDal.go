package dal

import "arvan-wallet/entity"

type UserTxnRepositoryInterface interface {
	SaveUserTxn(txn entity.UserTxn) error
	GetUserTxnsByUser(user entity.User) ([]entity.UserTxn, error)
}