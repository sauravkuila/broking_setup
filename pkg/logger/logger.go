package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sauravkuila/broking_setup/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logObject *zap.Logger

func Log(data ...context.Context) *zap.Logger {
	if data != nil {
		ctx := data[0]
		return logObject.With(zap.Any("requestID", ctx.Value(config.REQUESTID)), zap.Any("userID", ctx.Value(config.USERID)), zap.Any("ucc", ctx.Value(config.UCC)))
	} else {
		return logObject
	}
}

func LoggerInit(level zapcore.Level) {
	var (
		err error
	)
	fmt.Println("LOGGER INIT started")

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.FunctionKey = "f"
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	cfg.EncoderConfig.ConsoleSeparator = " "
	cfg.Encoding = "console"

	logObject, err = cfg.Build()
	if err != nil {
		fmt.Println("failed to create custom production logger , Exiting system", err)
		os.Exit(0)
	} else if logObject == nil {
		logObject, err = zap.NewProduction()
		logObject.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
		if err != nil {
			fmt.Println("failed to create production logger , Exiting system", err)
			os.Exit(0)
		}
		fmt.Println("Failed to create custom production logger, creating production logger")
	} else {
		logObject.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
		fmt.Println("custom production logger created")
	}

	Log().Info("Logger init successfully")
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}
