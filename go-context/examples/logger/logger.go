package logger

import (
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	defaultAppName   = "GLOBAL"
	defaultRequestID = "GLOBAL"
	defaultUserID    = 0
)

// START1 OMIT
type IRequestScopedLogger interface {
	ILogger

	GetRequestScoped(requestID, appName string, userID int64) ILogger
}

type ILogger interface {
	Printf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	*logrus.Logger
	requestID string
	appName   string
	userID    int64
}

// STOP1 OMIT

// START2 OMIT
func (l *logger) Printf(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"datetime":   time.Now(),
		"request_id": l.requestID,
		"app_name":   l.appName,
		"user_id":    l.userID,
	}).Infof(format, args...)
}

// STOP2 OMIT
func (l *logger) Errorf(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"datetime":   time.Now(),
		"request_id": l.requestID,
		"app_name":   l.appName,
		"user_id":    l.userID,
	}).Errorf(format, args...)
}

func (l *logger) GetRequestScoped(requestID, appName string, userID int64) ILogger {
	return &logger{l.Logger, requestID, appName, userID}
}

// START1A OMIT
func NewLogger() IRequestScopedLogger {
	l := logrus.New()

	return &logger{l, defaultAppName, defaultRequestID, defaultUserID}
}

// STOP1A OMIT
