package common

import (
	"code.google.com/p/log4go"
)

var Log log4go.Logger

func newLogger() log4go.Logger {
	logger := make(log4go.Logger)
	logger.LoadConfiguration("/opt/gainsboro/conf/gainsboro.xml")
	return logger
}

func init() {
	Log = newLogger()
	Log.Info("gainsboro service start")
}
