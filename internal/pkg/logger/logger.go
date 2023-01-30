package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates new instance of logger and sets it to global variable.
func New(level string) (*zap.Logger, error) {
	zaplevel, err := zapcore.ParseLevel(strings.ToUpper(level))
	if err != nil {
		return nil, err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zaplevel)

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(l)

	return l, nil
}
