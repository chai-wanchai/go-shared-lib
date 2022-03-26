package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLogger *zap.Logger

func New() error {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := loggerConfig.Build()
	if err != nil {
		return err
	}

	DefaultLogger = logger
	zap.ReplaceGlobals(logger)

	return nil
}
func Info(ctx context.Context, message string) {
	info(ctx, message)
}

func referenceID(ctx context.Context) zapcore.Field {
	return zap.String(HeaderReferenceID, GetFromContext(HeaderReferenceID, ctx))
}
func info(ctx context.Context, message string) {
	zap.L().With(
		referenceID(ctx),
		zap.String("log_level", "INFO"),
	).Info(message)
}
