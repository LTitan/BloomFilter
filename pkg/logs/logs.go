package logs

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger .
var Logger *zap.SugaredLogger

func initLogger() {
	writeSyncer := getLogWriter()
	encoder, console := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
		zapcore.NewCore(console, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getEncoder() (zapcore.Encoder, zapcore.Encoder) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig), zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/all.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func init() {
	initLogger()
}
