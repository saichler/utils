package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

func Fatal(any ...interface{}) {
	log.Fatal(any)
}

func Error(any ...interface{}) error {
	log.Error("******* ", any)
	return errors.New("*****")
}

func Warn(any ...interface{}) {
	log.Warn("***** ", any)
}

func Info(any ...interface{}) {
	log.Info("*** ", any)
}

func Debug(any ...interface{}) {
	log.Debug(any)
}

func Trace(any ...interface{}) {
	log.Trace(any)
}
