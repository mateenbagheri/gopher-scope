package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() error {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "level"

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(Logger)
	return nil
}

func SyncLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func WithFields(fields ...zap.Field) *zap.Logger {
	return Logger.With(fields...)
}

func GetLogger() *zap.Logger {
	return Logger
}

func RequestLogger(requestID, method, path string) *zap.Logger {
	return Logger.With(
		zap.String("request_id", requestID),
		zap.String("method", method),
		zap.String("path", path),
	)
}
