package tables

import "time"

type SubscriptionPlans struct {
	ID        int     `gorm:"column:subscription_plan_id;primaryKey;autoIncrement"`
	PID       string  `gorm:"column:subscription_plan_pid;unique;not null;type:varchar(40)"`
	Name      string  `gorm:"column:subscription_plan_name;not null;type:varchar(100)"`
	Price     float64 `gorm:"column:subscription_plan_price;not null;type:decimal(10,2)"`
	Coins     int     `gorm:"column:subscription_plan_coins;not null;type:int"`
	IsActive  bool    `gorm:"column:is_active;not null;default:true"`
	IsDeleted bool    `gorm:"column:is_deleted;not null;default:false"`
	IsSandbox bool    `gorm:"column:is_sandbox;not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
