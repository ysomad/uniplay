package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"go.uber.org/zap"
)

type statistic struct {
	log        *zap.Logger
	metricRepo metricRepository
}

func NewStatistic(l *zap.Logger, r metricRepository) *statistic {
	return &statistic{
		log:        l,
		metricRepo: r,
	}
}

func (s *statistic) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) (domain.WeaponStats, error) {
	metrics, err := s.metricRepo.GetWeaponMetrics(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	if len(metrics) == 0 {
		return nil, domain.ErrWeaponStatsNotFound
	}

	s.log.Error("test", zap.Any("test", metrics))

	stats := make(domain.WeaponStats)

	for _, m := range metrics {
		if _, ok := stats[m.WeaponName]; !ok {
			stats[m.WeaponName] = new(domain.WeaponStat)
		}

		ws, ok := stats[m.WeaponName]
		if ok {
			ws.SetStat(m.Metric, m.Value)
		}
	}

	return stats, nil
}

func (s *statistic) GetWeaponClassStats(ctx context.Context, steamID uint64) (domain.WeaponClassStats, error) {
	return nil, nil
}
