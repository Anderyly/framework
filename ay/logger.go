package ay

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	Logger *zap.Logger
)

func LoggerInit() {
	err := GetLogger(zapcore.DebugLevel) // Set the desired log level
	if err != nil {
		panic(err)
	}
	defer Logger.Sync()
}

func GetLogger(logLevel zapcore.Level) error {
	config := zap.Config{
		Encoding:    "json", // You can change this to "console" for human-readable output
		Level:       zap.NewAtomicLevelAt(logLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   getEncodingCfg,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}
	return nil
}

func getEncodingCfg(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000Z"))
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}
