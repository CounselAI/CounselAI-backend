package aisvc

import (
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type AiSvcImpl struct {
	usersGorm users.GormInterface
}

type Interface interface {
	GetCases(c *gin.Context) (utils.BaseResponse, error)
	Compile(c *gin.Context, reqBody CompileReq) (utils.BaseResponse, error)
	Query(c *gin.Context, reqBody QueryReq) (utils.BaseResponse, error)
}

func Handler(usersGorm users.GormInterface) Interface {
	return &AiSvcImpl{
		usersGorm: usersGorm,
	}
}
