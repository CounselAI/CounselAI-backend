package subscriptionplans

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GormInterface interface {
	GetSubscriptionPlanDetailsByPID(ctx *gin.Context, PID string) (tables.SubscriptionPlans, error)
	GetAllSubscriptionPlans(ctx *gin.Context) ([]tables.SubscriptionPlans, error)
}

func Gorm(gormDB *gorm.DB) *subscriptionPlansGormImpl {
	return &subscriptionPlansGormImpl{
		DB: gormDB,
	}
}

type subscriptionPlansGormImpl struct {
	DB *gorm.DB
}

func (r *subscriptionPlansGormImpl) GetSubscriptionPlanDetailsByPID(ctx *gin.Context, PID string) (tables.SubscriptionPlans, error) {
	var subscriptionPlans tables.SubscriptionPlans
	db := r.DB.Session(&gorm.Session{})

	err := db.Where("subscription_plan_pid = ?", PID).Find(&subscriptionPlans).Error
	if err != nil {
		return subscriptionPlans, errors.Wrap(err, "[subscriptionPlansGormImpl][GetSubscriptionPlanDetailsByPID]")
	}
	return subscriptionPlans, nil
}

func (r *subscriptionPlansGormImpl) GetAllSubscriptionPlans(ctx *gin.Context) ([]tables.SubscriptionPlans, error) {
	var subscriptionPlans []tables.SubscriptionPlans

	db := r.DB.Session(&gorm.Session{})

	result := db.Find(&subscriptionPlans).
		Where("is_active = ?", true).
		Where("is_deleted = ?", false).
		Order("subscription_plan_price ASC").
		Find(&subscriptionPlans)

	err := result.Error

	if err != nil {
		return subscriptionPlans, errors.Wrap(err, "[subscriptionPlansGormImpl][GetAllSubscriptionPlans]")
	}
	return subscriptionPlans, nil
}

func (r *subscriptionPlansGormImpl) CreateSubscriptionPlan(ctx *gin.Context, subscriptionPlans tables.SubscriptionPlans) (tables.SubscriptionPlans, error) {

	subscriptionPlans.PID = utils.UUIDWithPrefix(constants.Prefix.SUBSCRIPTIONSPLANS)
	db := r.DB.Session(&gorm.Session{})

	err := db.Create(&subscriptionPlans).Error
	if err != nil {
		return subscriptionPlans, errors.Wrap(err, "[subscriptionPlansGormImpl][CreateSubscriptionPlan]")
	}
	return subscriptionPlans, nil
}
