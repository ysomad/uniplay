package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
	"github.com/ssssargsian/uniplay/internal/postgres"
	"github.com/ssssargsian/uniplay/internal/service"
	"github.com/ysomad/pgxatomic"
)

func Run(conf *config.Config) {
	var err error

	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		log.Fatalf("pgclient.New: %s", err.Error())
	}

	pool, err := pgxatomic.NewPool(pgClient.Pool)
	if err != nil {
		log.Fatalf("pgxatomic.NewPool: %s", err.Error())
	}

	txrunner, err := pgxatomic.NewRunner(pgClient.Pool, pgx.TxOptions{})
	if err != nil {
		log.Fatalf("pgxatomic.NewRunner: %s", err.Error())
	}

	// repos
	matchRepo := postgres.NewMatchRepo(pool, pgClient.Builder)

	// services
	replayService := service.NewReplay(nil, matchRepo)

	// test
	replayFile, err := os.Open("./test-data/1.dem")
	if err != nil {
		log.Fatalf("open demo err: %s", err.Error())
	}
	defer replayFile.Close()

	var match *domain.Match

	// test run atomically
	err = txrunner.Run(context.Background(), func(txCtx context.Context) error {
		match, err = replayService.CollectStats(txCtx, replayFile)
		return err
	})
	if err != nil {
		log.Fatalf("demo collect error :%s", err.Error())
	}

	fmt.Println(match.ID.String())

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
