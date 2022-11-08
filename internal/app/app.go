package app

import (
	"log"
	"os"

	"github.com/ssssargsian/uniplay/internal/demoparser"
)

func Run() {
	demo, err := os.Open("./test-data/1.dem")
	if err != nil {
		log.Fatalf("failed to open demo file: %s", err.Error())
	}
	defer demo.Close()

	p := demoparser.New(demo)
	defer p.Close()

	// metrics, wmetrics, match, err := p.Parse()
	// if err != nil {
	// 	log.Fatalf("parse error: %s", err.Error())
	// }

	// metricsFile, err := json.MarshalIndent(metrics.Out(), "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("metrics.json", metricsFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }

	// wmetricsFile, err := json.MarshalIndent(wmetrics.Out(), "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("weapon_metrics.json", wmetricsFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }

	// matchFile, err := json.MarshalIndent(match, "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("match.json", matchFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }

}
