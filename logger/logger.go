package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	L      *logrus.Logger
	doOnce sync.Once
)

func Logger() *logrus.Logger {
	doOnce.Do(func() {
		L = newLogger()
	})
	return L
}

func newLogger() *logrus.Logger {
	// create new logger for the client
	//
	// https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#StandardLogger
	logger := logrus.StandardLogger()

	// assign contextual fields
	//
	// https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#NewEntry
	// logger = logrus.NewEntry(logger).WithField("engine", c.Driver())
	return logger
}
