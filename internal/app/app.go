package app

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"

	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
	"github.com/ssssargsian/uniplay/internal/postgres"
	"github.com/ssssargsian/uniplay/internal/service"
)

func Run(conf *config.Config) {
	var err error

	// db
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
	replayRepo := postgres.NewReplayRepo(pool, pgClient.Builder)

	// services
	replayService := service.NewReplay(replayRepo)

	// test
	replayFiles, err := os.ReadDir("./test-data/")
	if err != nil {
		log.Fatalf("ioutil.ReadDir: %s", err.Error())
	}

	for _, file := range replayFiles {
		if file.Name() == ".DS_Store" {
			continue
		}

		replayFile, err := os.Open("./test-data/" + file.Name())
		if err != nil {
			log.Fatalf("open file error: %s, replay filename %s", err.Error(), replayFile.Name())
		}
		defer replayFile.Close()

		err = txrunner.Run(context.Background(), func(txCtx context.Context) error {
			_, err = replayService.CollectStats(txCtx, replayFile)
			return err
		})

		if err != nil {
			log.Fatalf("demo collect error: %s, replay filename %s", err.Error(), replayFile.Name())
		}
	}
}
