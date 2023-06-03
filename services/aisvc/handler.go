package aisvc

import (
	"github.com/BearTS/go-gin-monolith/dbops/gorm/reports"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type AiSvcImpl struct {
	usersGorm   users.GormInterface
	reportsGorm reports.GormInterface
}

type Interface interface {
	GetCases(c *gin.Context) (utils.BaseResponse, error)
	Compile(c *gin.Context, reqBody CompileReq) (utils.BaseResponse, error)
	Query(c *gin.Context, reqBody QueryReq) (utils.BaseResponse, error)

	UploadReport(c *gin.Context) (utils.BaseResponse, error)
	GetAllReports(c *gin.Context) (utils.BaseResponse, error)
	ArchiveReport(c *gin.Context, reqBody ArchiveReportReq) (utils.BaseResponse, error)
}

func Handler(usersGorm users.GormInterface, reportsGorm reports.GormInterface) Interface {
	return &AiSvcImpl{
		usersGorm:   usersGorm,
		reportsGorm: reportsGorm,
	}
}
