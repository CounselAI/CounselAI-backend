package tables

import "time"

type Reports struct {
	ID         int    `gorm:"column:report_id;primaryKey;autoIncrement"`
	PID        string `gorm:"column:report_pid;unique;not null;type:varchar(40)"`
	Url        string `gorm:"column:url;not null;type:varchar(100)"`
	UserID     string `gorm:"column:user_id;not null;type:string"`
	IsSandbox  bool   `gorm:"column:is_sandbox;not null;default:false"`
	IsDeleted  bool   `gorm:"column:is_deleted;not null;default:false"`
	IsArchived bool   `gorm:"column:is_archived;not null;default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
