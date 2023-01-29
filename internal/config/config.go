package config

import "time"

type Config struct {
	App    App    `yaml:"app"`
	HTTP   HTTP   `yaml:"http"`
	Log    Log    `yaml:"log"`
	PG     PG     `yaml:"postgres"`
	Jaeger Jaeger `yaml:"jaeger"`
}

type App struct {
	Name        string `yaml:"name" env-required:"true"`
	Ver         string `yaml:"version" env-required:"true"`
	Environment string `yaml:"environment" env-required:"true"`
}

type (
	HTTP struct {
		Host string `yaml:"host" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
	}

	Log struct {
		Level   string         `yaml:"level" env-required:"true"`
		TimeLoc *time.Location `yaml:"time_location" env-default:"Etc/UTC"`
	}

	PG struct {
		MaxConns int32  `yaml:"max_connections" env-required:"true"`
		URL      string `env:"PG_URL" env-required:"true"`
	}

	Jaeger struct {
		Endpoint string `env:"JAEGER_ENDPOINT" env-required:"true"`
	}
)
