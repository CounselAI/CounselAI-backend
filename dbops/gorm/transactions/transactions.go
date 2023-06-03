package transactions

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GormInterface interface {
}

func Gorm(gormDB *gorm.DB) *transactionsGormImpl {
	return &transactionsGormImpl{
		DB: gormDB,
	}
}

type transactionsGormImpl struct {
	DB *gorm.DB
}

func (r *transactionsGormImpl) CreateTransaction(ctx *gin.Context, transactions tables.Transactions) (tables.Transactions, error) {

	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return transactions, errors.Wrap(err, "[transactionsGormImpl][CreateTransaction]")
	}

	transactions.UserPID = authData.UserPID
	transactions.PID = utils.UUIDWithPrefix(constants.Prefix.TRANSACTIONS)

	db := r.DB.Session(&gorm.Session{})

	err = db.Create(&transactions).Error
	if err != nil {
		return transactions, errors.Wrap(err, "[transactionsGormImpl][CreateTransaction]")
	}
	return transactions, nil
}

func (r *transactionsGormImpl) GetTransactionDetailsByPID(ctx *gin.Context, PID string) (tables.Transactions, error) {
	var transactions tables.Transactions

	db := r.DB.Session(&gorm.Session{})

	err := db.Where("transaction_pid = ?", PID).Find(&transactions).Error
	if err != nil {
		return transactions, errors.Wrap(err, "[transactionsGormImpl][GetTransactionDetailsByPID]")
	}
	return transactions, nil
}

func (r *transactionsGormImpl) GetAllTransactionsByUserPID(ctx *gin.Context, userPID string) ([]tables.Transactions, error) {
	var transactions []tables.Transactions

	db := r.DB.Session(&gorm.Session{})

	err := db.Where("user_pid = ?", userPID).Find(&transactions).Error
	if err != nil {
		return transactions, errors.Wrap(err, "[transactionsGormImpl][GetAllTransactionsByUserPID]")
	}
	return transactions, nil
}

// Update
func (r *transactionsGormImpl) UpdateTransaction(ctx *gin.Context, PID string, transactions tables.Transactions) (tables.Transactions, error) {
	var transaction tables.Transactions

	db := r.DB.Session(&gorm.Session{})

	err := db.Where("transaction_pid = ?", PID).
		Updates(transactions).Error
	if err != nil {
		return transaction, errors.Wrap(err, "[transactionsGormImpl][UpdateTransaction]")
	}

	return transaction, nil
}
