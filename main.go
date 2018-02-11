package main

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// DecorateRuntimeContext appends line, file and function context to the logger
func DecorateRuntimeContext(logger *logrus.Entry) *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", file).WithField("line", line).WithField("func", fName)
	} else {
		return logger
	}
}

func initLogger() {
	// logger = DecorateRuntimeContext
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
}
func main() {
	initLogger()
	bc := NewBlockchain()
	defer bc.db.Close()
	cli := CLI{bc}
	cli.Run()
}
