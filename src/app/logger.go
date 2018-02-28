package app

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var Log *logrus.Logger

func InitLog() {
	//log.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	Log = logrus.New()
}
