package raptor

import (
	"os"

	"github.com/sirupsen/logrus"
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
