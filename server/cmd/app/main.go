package main

import (
	"flag"
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ysomad/uniplay/internal/app"
	"github.com/ysomad/uniplay/internal/config"
)

func main() {
	var flags app.Flags

	flag.BoolVar(&flags.Migrate, "migrate", false, "run migrations on start")
	flag.StringVar(&flags.MigrationsDir, "migrations-dir", "./migrations", "path to migrations directory")
	flag.StringVar(&flags.ConfigPath, "conf", "./configs/local.yml", "path to yml config")
	flag.Parse()

	var conf config.Config

	if err := cleanenv.ReadConfig(flags.ConfigPath, &conf); err != nil {
		log.Fatalf("config parse error: %s", err)
	}

	app.Run(&conf, flags)
}
