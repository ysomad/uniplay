package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ysomad/uniplay/internal/app"
	"github.com/ysomad/uniplay/internal/config"
)

func main() {
	var conf config.Config

	// TODO: parse config from flag or ENVIRONMENT VARIABLE
	if err := cleanenv.ReadConfig("./configs/local.yml", &conf); err != nil {
		log.Fatalf("config parse error: %s", err)
	}

	app.Run(&conf)
}
