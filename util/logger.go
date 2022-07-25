package util

import (
	"encoding/json"
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

	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		panic(err)
	}

	wInfo := zapcore.AddSync(rotatorInfo)
	wError := zapcore.AddSync(rotatorError)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), wInfo, zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), wError, zapcore.ErrorLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
