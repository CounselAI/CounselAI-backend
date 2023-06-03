package ai

import (
	"github.com/BearTS/go-gin-monolith/services/aisvc"
)

type AiHandler struct {
	aiSvc aisvc.Interface
}

func Handler(aiSvc aisvc.Interface) *AiHandler {
	return &AiHandler{
		aiSvc: aiSvc,
	}
}
