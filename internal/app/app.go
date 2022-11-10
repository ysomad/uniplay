package app

import (
	"github.com/ssssargsian/uniplay/internal/config"
)

func Run(conf *config.Config) {

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
