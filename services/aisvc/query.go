package aisvc

import (
	"net/http"

	counselai "github.com/BearTS/go-gin-monolith/providers/counsel_ai"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *AiSvcImpl) Query(c *gin.Context, reqBody QueryReq) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// Check if user has enough coin
	user, err := g.usersGorm.GetUserDetails(c)

	if user.AvailableCoins < 1 {
		baseRes.Success = false
		baseRes.StatusCode = http.StatusPaymentRequired
		baseRes.Message = "Not enough coins"
		return baseRes, err
	}

	// Get Cases Data from Provider
	var req counselai.QueryReq
	req.Query = reqBody.Query
	dataRes, err := counselai.Query(c, req)
	if err != nil {
		return baseRes, err
	}

	if !dataRes.Success {
		baseRes.Success = dataRes.Success
		baseRes.StatusCode = dataRes.StatusCode
		baseRes.Message = dataRes.Message
		baseRes.Data = dataRes.Data

		return baseRes, err
	}

	baseRes.Success = dataRes.Success
	baseRes.StatusCode = dataRes.StatusCode
	baseRes.Message = dataRes.Message
	baseRes.Data = dataRes.Data

	return baseRes, err
}
