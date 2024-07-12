package config

import (
	"davinci-chat/consts"
	"fmt"
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

func GetDataSourceName(dbName string) string {
	runEnv := GetRunEnv()

	switch runEnv {
	case consts.Production:
		dsName := fmt.Sprintf("/var/chat/%s.sqlite", dbName)
		return dsName
	case consts.Development:
		dsName := fmt.Sprintf("./%s.sqlite", dbName)
		return dsName
	default:
		log.Fatalf("Invalid RUN_ENV value: %s", runEnv)
		return ""
	}
}
