package subscriptionplansvc

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (s *subscriptionplanSvcImpl) GetAllSubscriptionPlans(ctx *gin.Context) (utils.BaseResponse, []tables.SubscriptionPlans, error) {
	var baseRes utils.BaseResponse
	var res []tables.SubscriptionPlans
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	res, err = s.subscriptionplansGorm.GetAllSubscriptionPlans(ctx)

	if err != nil {
		return baseRes, res, err
	}

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Success"

	return baseRes, res, err
}
