package aisvc

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *AiSvcImpl) GetAllReports(c *gin.Context) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	archivedFitler := c.Query("archived")
	var reports []tables.Reports
	switch archivedFitler {
	case "true":
		reports, err = g.reportsGorm.GetAllReportsWithFilters(c, true)
	case "false":
		reports, err = g.reportsGorm.GetAllReportsWithFilters(c, false)
	default:
		reports, err = g.reportsGorm.GetAllReports(c)
	}

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Success"
	baseRes.Data = reports

	return baseRes, err

}
