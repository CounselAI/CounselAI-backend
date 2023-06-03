package subscriptionplansvc

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	subscriptionplans "github.com/BearTS/go-gin-monolith/dbops/gorm/subscription_plans"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type subscriptionplanSvcImpl struct {
	subscriptionplansGorm subscriptionplans.GormInterface
}

// interface.
type Interface interface {
	GetAllSubscriptionPlans(ctx *gin.Context) (utils.BaseResponse, []tables.SubscriptionPlans, error)
}

func Handler(subscriptionplansGorm subscriptionplans.GormInterface) Interface {
	return &subscriptionplanSvcImpl{
		subscriptionplansGorm: subscriptionplansGorm,
	}
}
