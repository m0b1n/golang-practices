package common

import (
	"arvan-wallet/dal/connection"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

//StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

var mutex = &sync.Mutex{}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := connection.GetPostgresInstance(connection.GetPostgresConfig())
		mutex.Lock()
		txHandle := db.Begin()

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			txHandle.Rollback()
		}
		mutex.Unlock()
	}
}
