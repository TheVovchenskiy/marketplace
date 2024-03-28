package utils

import (
	"context"
	"marketplace/pkg/token"
	"strconv"

	"github.com/sirupsen/logrus"
)

type ContextKey string

const (
	REQUEST_ID_KEY = ContextKey("request_id")
	LOGGER_KEY     = ContextKey("logger")
	USER_ID_KEY    = ContextKey("user_id")
)

func GetContextUserId(ctx context.Context) (int, error) {
	userIdCtx, ok := ctx.Value(USER_ID_KEY).(string)
	if !ok {
		return 0, ErrNoUserIdInContext
	}
	userId, err := strconv.Atoi(userIdCtx)
	if err != nil {
		return 0, token.ErrInvalidToken
	}
	return userId, nil
}

func GetContextLogger(ctx context.Context) *logrus.Entry {
	logger, ok := ctx.Value(LOGGER_KEY).(*logrus.Entry)
	if !ok {
		defaultLogger := logrus.New()
		defaultLogger.SetLevel(logrus.InfoLevel)
		return defaultLogger.WithField("default", true)
	}
	return logger
}

func GetRequestIDFromCtx(ctx context.Context) string {
	return ctx.Value(REQUEST_ID_KEY).(string)
}

func UpdateCtxLoggerWithMethod(ctx context.Context, methodName string) context.Context {
	contextLogger := GetContextLogger(ctx)
	newContextLogger := contextLogger.WithFields(logrus.Fields{
		"method": methodName,
	})
	return context.WithValue(ctx, LOGGER_KEY, newContextLogger)
}
