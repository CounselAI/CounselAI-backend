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
	GetAllReportsWithFilters(ctx *gin.Context, isArchived bool) ([]tables.Reports, error)
	UpdateReport(ctx *gin.Context, report tables.Reports) (tables.Reports, error)
	GetReportByPID(ctx *gin.Context, reportPID string) (tables.Reports, error)
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

func (r *reportGormImpl) GetReportByPID(ctx *gin.Context, reportPID string) (tables.Reports, error) {
	var report tables.Reports

	err := r.DB.Session(&gorm.Session{}).Where("reports_pid = ?", reportPID).Find(&report).Error
	if err != nil {
		return report, err
	}

	return report, nil
}

func (r *reportGormImpl) GetAllReportsWithFilters(ctx *gin.Context, isArchived bool) ([]tables.Reports, error) {
	var reports []tables.Reports
	authData, _ := utils.GetAuthData(ctx)

	err := r.DB.Session(&gorm.Session{}).Where("user_id = ? AND is_archived = ?", authData.UserPID, isArchived).Find(&reports).Error
	if err != nil {
		return reports, err
	}

	return reports, nil
}

//  update report

func (r *reportGormImpl) UpdateReport(ctx *gin.Context, report tables.Reports) (tables.Reports, error) {

	db := r.DB.Session(&gorm.Session{})

	err := db.Where("reports_pid = ?", report.PID).Updates(&report).Error
	if err != nil {
		return report, err
	}

	return report, nil
}
