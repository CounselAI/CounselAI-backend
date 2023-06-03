package aisvc

import (
	"net/http"

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

	// Get All Reports Data from Provider
	reports, err := g.reportsGorm.GetAllReports(c)

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Success"
	baseRes.Data = reports

	return baseRes, err

}
