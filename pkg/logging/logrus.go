package logging

import (
	"marketplace/configs"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger(logFile *os.File) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: "Mon, 02 Jan 2006 15:04:05 MST",
	}

	Logger.SetOutput(logFile)
	Logger.SetFormatter(jsonFormatter)
	Logger.SetLevel(configs.LOG_LEVEL)
}
