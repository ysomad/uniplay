package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/service"
)

func Run(conf *config.Config) {

	replayService := service.NewReplay(nil, nil)

	replayFile, err := os.Open("./test-data/1.dem")
	if err != nil {
		log.Fatalf("open demo err: %s", err.Error())
	}
	defer replayFile.Close()

	m, err := replayService.CollectStats(context.Background(), replayFile)
	if err != nil {
		log.Fatalf("collect error :%s", err.Error())
	}

	fmt.Println(m.ID.String())

	// metricsFile, err := json.MarshalIndent(res.Metrics.ToDTO(uuid.UUID{}), "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("metrics_dto.json", metricsFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }

	// wmetricsFile, err := json.MarshalIndent(res.WeaponMetrics.ToDTO(uuid.UUID{}), "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("weapon_metrics_dto.json", wmetricsFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }

	// matchFile, err := json.MarshalIndent(res.Match, "", " ")
	// if err != nil {
	// 	log.Fatalf("json.MarshalIndent: %s", err.Error())
	// }

	// if err = os.WriteFile("match.json", matchFile, 0644); err != nil {
	// 	log.Fatalf("ioutil.WriteFile: %s", err.Error())
	// }
}
