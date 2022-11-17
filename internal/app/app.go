package app

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/pkg/logger"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
	"github.com/ssssargsian/uniplay/internal/postgres"
	"github.com/ssssargsian/uniplay/internal/service"
)

func Run(conf *config.Config) {
	var err error

	l, err := logger.New(os.Stderr, conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	// db
	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	pool, err := pgxatomic.NewPool(pgClient.Pool)
	if err != nil {
		l.Fatal("pgxatomic.NewPool", zap.Error(err))
	}

	atomic, err := pgxatomic.NewRunner(pgClient.Pool, pgx.TxOptions{})
	if err != nil {
		l.Fatal("pgxatomic.NewRunner", zap.Error(err))
	}

	// repos
	replayRepo := postgres.NewReplayRepo(pool, pgClient.Builder)

	// services
	replayService := service.NewReplay(l, replayRepo)

	// test all
	replayFiles, err := os.ReadDir("./test-data/")
	if err != nil {
		l.Fatal("os.ReadDir", zap.Error(err))
	}

	for _, file := range replayFiles {
		if file.Name() == ".DS_Store" {
			continue
		}

		replayFile, err := os.Open("./test-data/" + file.Name())
		if err != nil {
			l.Fatal("open file error", zap.Error(err))
		}
		defer replayFile.Close()

		err = atomic.Run(context.Background(), func(txCtx context.Context) error {
			_, err = replayService.CollectStats(txCtx, replayFile)
			return err
		})
		if err != nil {
			l.Fatal("demo collect error", zap.Error(err))
		}
	}

	// test one
	// replayFile, err := os.Open("./test-data/5.dem")
	// if err != nil {
	// 	l.Fatal("open file error", zap.Error(err))
	// }
	// defer replayFile.Close()

	// err = atomic.Run(context.Background(), func(txCtx context.Context) error {
	// 	_, err = replayService.CollectStats(txCtx, replayFile)
	// 	return err
	// })
	// if err != nil {
	// 	l.Fatal("demo collect error", zap.Error(err))
	// }

}
