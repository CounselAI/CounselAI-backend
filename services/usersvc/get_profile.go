package usersvc

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *UserSvcImpl) GetProfile(c *gin.Context) (utils.BaseResponse, tables.Users, error) {
	var baseRes utils.BaseResponse
	var res tables.Users
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// search user by PID
	authData, err := utils.GetAuthData(c)
	if err != nil {
		baseRes.StatusCode = http.StatusUnauthorized
		baseRes.Message = "Unauthorized"
		return baseRes, res, err
	}

	res, err = g.usersGorm.GetUserDetailsByPID(c, authData.UserPID)
	if err != nil {
		return baseRes, res, err
	}

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Success"

	return baseRes, res, err
}
