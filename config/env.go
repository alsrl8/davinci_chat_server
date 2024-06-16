package config

import (
	"davinci-chat/consts"
	"log"
	"os"
)

func GetRunEnv() consts.RunEnv {
	runEnv := consts.RunEnv(os.Getenv("RUN_ENV"))

	switch runEnv {
	case consts.Development, consts.Production:
		return runEnv
	default:
		log.Fatalf("Invalid RUN_ENV value: %s", runEnv)
		return ""
	}
}
