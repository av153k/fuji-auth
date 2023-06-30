package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	flavor := os.Getenv("FLAVOR")
	if flavor == "DEV" {
		log.SetLevel(log.DebugLevel)
	}
}
