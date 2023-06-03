package ai

import (
	"github.com/BearTS/go-gin-monolith/services/aisvc"
	"github.com/gin-gonic/gin"
)

func validateCompileReq(c *gin.Context) (aisvc.CompileReq, error) {
	var reqBody aisvc.CompileReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}

	return reqBody, nil
}

func validateQueryReq(c *gin.Context) (aisvc.QueryReq, error) {
	var reqBody aisvc.QueryReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}

	return reqBody, nil
}
