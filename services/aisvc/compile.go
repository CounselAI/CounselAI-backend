package aisvc

import (
	"fmt"
	"net/http"

	counselai "github.com/BearTS/go-gin-monolith/providers/counsel_ai"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *AiSvcImpl) Compile(c *gin.Context, reqBody CompileReq) (utils.BaseResponse, error) {
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

	// Check if all the ids and coins are in the correct format
	if len(reqBody.IDs) > user.AvailableCoins {
		baseRes.Success = false
		baseRes.StatusCode = http.StatusPaymentRequired
		baseRes.Message = "Not enough coins"
		return baseRes, err
	}

	// Get Cases Data from Provider
	var req counselai.CompileReq
	req.Ids = reqBody.IDs
	dataRes, err := counselai.Compile(c, req, reqBody.TypeOfAi)
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

	// Update user's available coins
	user.AvailableCoins = user.AvailableCoins - len(reqBody.IDs)
	user, err = g.usersGorm.UpdateUser(c, user)
	if err != nil {
		fmt.Println(err)
	}

	baseRes.Success = dataRes.Success
	baseRes.StatusCode = dataRes.StatusCode
	baseRes.Message = dataRes.Message
	baseRes.Data = dataRes.Data

	return baseRes, err
}
