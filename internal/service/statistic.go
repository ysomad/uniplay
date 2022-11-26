package service

import (
	"context"
	"sort"

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

func (s *statistic) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error) {
	rawMetrics, err := s.metricRepo.GetWeaponMetrics(ctx, steamID, f)
	if err != nil {
		return nil, err
	}
	if len(rawMetrics) == 0 {
		return nil, domain.ErrWeaponStatsNotFound
	}

	var out []domain.WeaponStats

	for _, m := range rawMetrics {
		i := sort.Search(len(out), func(i int) bool {
			return out[i].Weapon == m.Weapon
		})

		if i < len(out) && out[i].Weapon == m.Weapon {
			out[i].Stats.SetStat(m.Metric, m.Value)
			continue
		}

		s := new(domain.WeaponStat)
		s.SetStat(m.Metric, m.Value)

		out = append(out, domain.WeaponStats{
			WeaponID: m.WeaponID,
			Weapon:   m.Weapon,
			ClassID:  m.ClassID,
			Class:    m.Class,
			Stats:    s,
		})
	}

	return out, nil
}

// func (s *statistic) GetWeaponClassStats(ctx context.Context, steamID uint64, c domain.WeaponClassID) (domain.WeaponClassStats, error) {
// metrics, err := s.metricRepo.GetWeaponClassMetrics(ctx, steamID, c)
// if err != nil {
// 	return nil, err
// }

// if len(metrics) == 0 {
// 	return nil, domain.ErrWeaponClassStatsNotFound
// }

// // TODO: refactor
// stats := make(domain.WeaponClassStats)
// for _, m := range metrics {
// 	wc := m.WeaponClassID.String()
// 	if _, ok := stats[wc]; !ok {
// 		stats[wc] = new(domain.WeaponStat)
// 	}

// 	ws, ok := stats[wc]
// 	if ok {
// 		ws.SetStat(m.Metric, m.Value)
// 	}
// }

// return stats, nil
// }
