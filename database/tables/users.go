package tables

import (
	"time"
)

type Users struct {
	ID             int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID            string `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	Email          string `gorm:"column:user_email;type:varchar(100)"`
	Name           string `gorm:"column:user_name;not null;type:varchar(100)"`
	MobileNumber   string `gorm:"column:user_mobile_number;not null;type:varchar(10)"`
	AvailableCoins int    `gorm:"column:user_available_coins;not null;type:int;default:4"`
	Metadata       JSONB  `gorm:"column:metadata;type:json"`
	IsDeleted      bool   `gorm:"column:is_deleted;not null;default:false"`
	IsSandbox      bool   `gorm:"column:is_sandbox;not null;default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
