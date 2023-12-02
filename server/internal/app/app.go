package app

import (
	"log/slog"
	"os"

	"github.com/ysomad/uniplay/internal/config"
)

func Run(conf *config.Config) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	})))

	runHTTPServer(conf)
}
