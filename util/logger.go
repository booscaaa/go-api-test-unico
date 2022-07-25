package util

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitializeLogger() *zap.Logger {
	logFileInfo := "/var/log/info-api-test-unico-%Y-%m-%d-%H.log"
	rotatorInfo, err := rotatelogs.New(
		logFileInfo,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		panic(err)
	}

	logFileError := "/var/log/error-api-test-unico-%Y-%m-%d-%H.log"
	rotatorError, err := rotatelogs.New(
		logFileError,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	wInfo := zapcore.AddSync(rotatorInfo)
	wError := zapcore.AddSync(rotatorError)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), wInfo, zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), wError, zapcore.ErrorLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
