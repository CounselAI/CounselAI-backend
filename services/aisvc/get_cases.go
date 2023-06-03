package aisvc

import (
	"net/http"

	counselai "github.com/BearTS/go-gin-monolith/providers/counsel_ai"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *AiSvcImpl) GetCases(c *gin.Context) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// Get Cases Data from Provider
	dataRes, err := counselai.GetCases(c)

	baseRes.Success = dataRes.Success
	baseRes.StatusCode = dataRes.StatusCode
	baseRes.Message = dataRes.Message
	baseRes.Data = dataRes.Data

	return baseRes, err

}
