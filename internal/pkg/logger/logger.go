package logger

import (
	"errors"
	"io"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var errNilWriter = errors.New("nil writer")

func New(w io.Writer, level string) (*zap.Logger, error) {
	if w == nil {
		return nil, errNilWriter
	}

	conf := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf.EncoderConfig),
		zapcore.AddSync(w),
		parseLevel(level),
	)

	return zap.New(core), nil
}

func parseLevel(l string) zapcore.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
