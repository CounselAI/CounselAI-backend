package aisvc

import (
	"errors"
	"net/http"

	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (g *AiSvcImpl) ArchiveReport(c *gin.Context, reqBody ArchiveReportReq) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// Get Report By ID
	report, err := g.reportsGorm.GetReportByPID(c, reqBody.PID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseRes.Success = false
			baseRes.StatusCode = http.StatusNotFound
			baseRes.Message = "Report not found"
			return baseRes, err
		}
		return baseRes, err
	}

	report.IsArchived = true

	// Update
	report, err = g.reportsGorm.UpdateReport(c, report)
	if err != nil {
		return baseRes, err
	}

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Report archived successfully"
	baseRes.Data = report

	return baseRes, err
}
