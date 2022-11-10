package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ssssargsian/uniplay/internal/app"
	"github.com/ssssargsian/uniplay/internal/config"
)

func main() {
	var conf config.Config
	if err := cleanenv.ReadConfig("./configs/local.yml", &conf); err != nil {
		log.Fatalf("config parse error: %s", err)
	}

	app.Run(&conf)
}
