package config

import (
	"log/slog"
	"strings"
	"time"
)

type Config struct {
	App           App           `toml:"app"`
	HTTP          HTTP          `toml:"http"`
	Log           Log           `toml:"log"`
	ObjectStorage ObjectStorage `toml:"object_storage"`
	PG            PG            `toml:"postgres"`
	Kratos        Kratos        `toml:"kratos"`
	Connect       Connect       `toml:"connect"`
}

type (
	App struct {
		Name        string `toml:"name" env-required:"true"`
		Ver         string `toml:"version" env-required:"true"`
		Environment string `toml:"environment" env-required:"true"`
	}

	HTTP struct {
		Host string `toml:"host" env-required:"true"`
		Port string `toml:"port" env-required:"true"`
	}

	Connect struct {
		Host string `toml:"host" env-required:"true"`
		Port string `toml:"port" env-required:"true"`
	}

	PG struct {
		URL      string `env:"PG_URL" env-required:"true"`
		DB       string `env:"PG_DB" env-required:"true"`
		MaxConns int32  `toml:"max_connections" env-required:"true"`
	}

	Kratos struct {
		URL               string        `toml:"url" env:"KRATOS_PUBLIC_URL" env-required:"true"`
		OrganizerSchemaID string        `toml:"organizer_schema_id" env-required:"true"`
		ClientTimeout     time.Duration `toml:"client_timeout" env-required:"true"`
		Debug             bool          `toml:"debug"`
	}

	ObjectStorage struct {
		Endpoint   string `toml:"endpoint" env-required:"true"`
		AccessKey  string `env:"S3_ACCESS_KEY" env-required:"true"`
		SecretKey  string `env:"S3_SECRET_KEY" env-required:"true"`
		DemoBucket string `toml:"demo_bucket" env-required:"true"`
		SSL        bool   `toml:"ssl"`
	}
)

type Log struct {
	Level string `toml:"level" env-required:"true"`
}

func (l *Log) SlogLevel() slog.Level {
	switch strings.ToLower(l.Level) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error", "err":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
