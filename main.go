package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// DecorateRuntimeContext appends line, file and function context to the logger
func DecorateRuntimeContext(logger *logrus.Entry) *logrus.Entry {
	if _, file, line, ok := runtime.Caller(1); ok {
		file := strings.Split(file, "/")
		return logger.WithField("file", file[len(file)-1]).WithField("line", line)
	} else {
		return logger
	}
}

func initLogger() {
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	// log.SetOutput(&lumberjack.Logger{
	// 	Filename:   "./foo.log",
	// 	MaxSize:    1, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     28,   //days
	// 	Compress:   true, // disabled by default
	// })
}
func main() {
	initLogger()
	DecorateRuntimeContext(logrus.NewEntry(logger)).Info("Test file name")
	cli := CLI{}
	cli.Run()
}
