package tables

import (
	"time"
)

type Transactions struct {
	ID                 int       `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	PID                string    `gorm:"column:transaction_pid;unique;not null;type:varchar(40)"`
	UserPID            string    `gorm:"column:user_pid;not null;type:varchar(40)"`
	// ProviderPID        string    `gorm:"column:provider_pid;not null;type:varchar(40)"`
	SubscriptionPlanID int       `gorm:"column:subscription_plan_id;not null;type:int"`
	Amount             float64   `gorm:"column:transaction_amount;not null;type:decimal(10,2)"`
	Status             string    `gorm:"column:transaction_status;not null;type:varchar(20)"`
	IsDeleted          bool      `gorm:"column:is_deleted;not null;default:false"`
	IsSandbox          bool      `gorm:"column:is_sandbox;not null;default:false"`
	CreatedAt          time.Time `gorm:"column:created_at"`
}
