package ai

import (
	"github.com/BearTS/go-gin-monolith/merrors"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (h *AiHandler) GetCases(c *gin.Context) {
	baseRes, _ := h.aiSvc.GetCases(c)

	utils.ReturnJSONStruct(c, baseRes)
}

func (h *AiHandler) Compile(c *gin.Context) {
	reqBody, err := validateCompileReq(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}
	baseRes, _ := h.aiSvc.Compile(c, reqBody)

	utils.ReturnJSONStruct(c, baseRes)
}

func (h *AiHandler) Query(c *gin.Context) {
	reqBody, err := validateQueryReq(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}
	baseRes, _ := h.aiSvc.Query(c, reqBody)

	utils.ReturnJSONStruct(c, baseRes)
}

func (h *AiHandler) UploadReport(c *gin.Context) {
	baseRes, _ := h.aiSvc.UploadReport(c)

	utils.ReturnJSONStruct(c, baseRes)
}

func (h *AiHandler) GetAllReports(c *gin.Context) {
	baseRes, _ := h.aiSvc.GetAllReports(c)

	utils.ReturnJSONStruct(c, baseRes)
}

func (h *AiHandler) ArchiveReport(c *gin.Context) {
	reqBody, err := validareArchiveReport(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}

	baseRes, _ := h.aiSvc.ArchiveReport(c, reqBody)

	utils.ReturnJSONStruct(c, baseRes)
}
