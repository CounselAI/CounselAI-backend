package reports

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormInterface interface {
	GetAllReports(ctx *gin.Context) ([]tables.Reports, error)
	CreateNewReport(ctx *gin.Context, report tables.Reports) (tables.Reports, error)
}

type reportGormImpl struct {
	DB *gorm.DB
}

func Gorm(gormDB *gorm.DB) *reportGormImpl {
	return &reportGormImpl{
		DB: gormDB,
	}
}

func (r *reportGormImpl) GetAllReports(ctx *gin.Context) ([]tables.Reports, error) {
	var reports []tables.Reports
	authData, _ := utils.GetAuthData(ctx)

	err := r.DB.Session(&gorm.Session{}).Where("user_id = ?", authData.UserPID).Find(&reports).Error
	if err != nil {
		return reports, err
	}

	return reports, nil
}

func (r *reportGormImpl) CreateNewReport(ctx *gin.Context, report tables.Reports) (tables.Reports, error) {
	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return report, err
	}

	report.UserID = authData.UserPID
	report.PID = utils.UUIDWithPrefix(constants.Prefix.REPORT)

	err = r.DB.Session(&gorm.Session{}).Create(&report).Error
	if err != nil {
		return report, err
	}
	return report, nil
}
