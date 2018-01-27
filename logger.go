package raptor

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var logger = logrus.New()

func getLogger() *logrus.Logger {
	return logger
}

func init() {
	if os.Getenv("RAPTOR_LOG") != "" {
		getLogger().SetLevel(logrus.DebugLevel)
	}
}
