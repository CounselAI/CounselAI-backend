package transactions

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *transactionsSvcImpl) CreateTransaction(ctx *gin.Context, reqBody CreateTransactionReq) (utils.BaseResponse, tables.Transactions, error) {
	panic("implement me")
}
