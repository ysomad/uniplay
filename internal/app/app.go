package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ssssargsian/uniplay/demoparser"
)

func Run() {
	demo, err := os.Open("./test-data/1.dem")
	if err != nil {
		log.Fatalf("failed to open demo file: %s", err.Error())
	}
	defer demo.Close()

	p := demoparser.New(demo)
	defer p.Close()

	metrics, weaponEvents, roundScore, err := p.Parse()
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}

	fmt.Println(metrics)
	fmt.Println(weaponEvents)
	fmt.Println(roundScore)
}
