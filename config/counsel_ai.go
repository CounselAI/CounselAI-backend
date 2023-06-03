package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type CounselAiConfig struct {
	URL string `split_words:"true" json:"COUNSEL_AI_URL"`
}

var CounselAi *CounselAiConfig

func loadCounselAiConfig() {
	CounselAi = &CounselAiConfig{}
	err := envconfig.Process("counsel_ai", CounselAi)
	if err != nil {
		log.Fatal(err.Error())
	}
}
