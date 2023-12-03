package config

import (
	"log/slog"
	"strings"
	"time"
)

type Config struct {
	App           App           `yaml:"app"`
	HTTP          HTTP          `yaml:"http"`
	Log           Log           `yaml:"log"`
	ObjectStorage ObjectStorage `yaml:"object_storage"`
	PG            PG            `yaml:"postgres"`
	Kratos        Kratos        `yaml:"kratos"`
	Connect       Connect       `yaml:"connect"`
}

type (
	App struct {
		Name        string `yaml:"name" env-required:"true"`
		Ver         string `yaml:"version" env-required:"true"`
		Environment string `yaml:"environment" env-required:"true"`
	}

	HTTP struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
	}

	Connect struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
	}

	PG struct {
		URL      string `env:"PG_URL" env-required:"true"`
		DB       string `env:"PG_DB" env-required:"true"`
		MaxConns int32  `yaml:"max_connections" env-required:"true"`
	}

	Kratos struct {
		URL               string        `yaml:"url" env:"KRATOS_PUBLIC_URL" env-required:"true"`
		OrganizerSchemaID string        `yaml:"organizer_schema_id" env-required:"true"`
		ClientTimeout     time.Duration `yaml:"client_timeout" env-required:"true"`
		Debug             bool          `yaml:"debug"`
	}

	ObjectStorage struct {
		Endpoint   string `yaml:"endpoint" env-required:"true"`
		AccessKey  string `env:"S3_ACCESS_KEY" env-required:"true"`
		SecretKey  string `env:"S3_SECRET_KEY" env-required:"true"`
		DemoBucket string `yaml:"demo_bucket" env-required:"true"`
		SSL        bool   `yaml:"ssl"`
	}
)

type Log struct {
	Level string `yaml:"level" env-required:"true"`
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
