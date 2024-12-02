package logger

import (
	"os"
	"payments-service/pkg/constants"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(debug bool) *logrus.Logger {
	log = logrus.New()

	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(logrus.DebugLevel)
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	return log
}

func Info(message string, category string) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Info(message)
}

func InfoF(format string, category string, args ...interface{}) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Infof(format, args...)
}

func Debug(message string, category string) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Debug(message)
}

func DebugF(format string, category string, args ...interface{}) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Debugf(format, args...)
}

func Error(message string, category string) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Error(message)
}

func ErrorF(format string, category string, args ...interface{}) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Errorf(format, args...)
}

func Fatal(message string, category string) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Fatal(message)
}

func FatalF(format string, category string, args ...interface{}) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Fatalf(format, args...)
}

func Panic(message string, category string) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Panic(message)
}

func PanicF(format string, category string, args ...interface{}) {
	fields := logrus.Fields{constants.LoggerCategory: category}

	log.WithFields(fields).Panicf(format, args...)
}
