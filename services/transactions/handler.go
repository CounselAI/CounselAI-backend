package transactions

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	subscriptionplans "github.com/BearTS/go-gin-monolith/dbops/gorm/subscription_plans"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/transactions"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type transactionsSvcImpl struct {
	transactionsGorm      transactions.GormInterface
	usersGorm             users.GormInterface
	subscriptionplansGorm subscriptionplans.GormInterface
}

type Interface interface {
	CreateTransaction(ctx *gin.Context, reqBody CreateTransactionReq) (utils.BaseResponse, tables.Transactions, error)
	// To do Update transaction
}

func Handler(transactionsGorm transactions.GormInterface, usersGorm users.GormInterface, subscriptionplansGorm subscriptionplans.GormInterface) Interface {
	return &transactionsSvcImpl{
		transactionsGorm:      transactionsGorm,
		usersGorm:             usersGorm,
		subscriptionplansGorm: subscriptionplansGorm,
	}
}
