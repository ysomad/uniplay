package service

import (
	"context"
	"sort"

	"github.com/ssssargsian/uniplay/internal/domain"
	"go.uber.org/zap"
)

type statistic struct {
	log  *zap.Logger
	repo statisticRepository
}

func NewStatistic(l *zap.Logger, r statisticRepository) *statistic {
	return &statistic{
		log:  l,
		repo: r,
	}
}

func (s *statistic) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) (map[uint16]domain.WeaponStats, error) {
	rawMetrics, err := s.repo.GetWeaponStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	if len(rawMetrics) == 0 {
		return nil, domain.ErrWeaponStatsNotFound
	}

	var (
		weaponStats = make(map[uint16]domain.WeaponStats)
		// out         []domain.WeaponStats
	)

	for _, m := range rawMetrics {
		if s, ok := weaponStats[m.WeaponID]; ok {
			s.SetStats(m.Metric, m.Value)
			continue
		}

		stats := domain.WeaponStats{
			WeaponID:      m.WeaponID,
			Weapon:        m.Weapon,
			Stats:         new(domain.WeaponStat),
			AccuracyStats: new(domain.WeaponAccuracyStat),
		}
		stats.SetStats(m.Metric, m.Value)

		weaponStats[m.WeaponID] = stats
	}

	s.log.Error("YEET", zap.Any("YEET", weaponStats))

	// for _, m := range rawMetrics {
	// 	i := sort.Search(len(out), func(i int) bool {
	// 		return out[i].Weapon == m.Weapon
	// 	})

	// 	if i < len(out) && out[i].Weapon == m.Weapon {
	// 		out[i].SetStat(m.Metric, m.Value)
	// 		continue
	// 	}

	// 	weaponStat := new(domain.WeaponStat)
	// 	weaponStat.SetStat(m.Metric, m.Value)

	// 	accuracyStat := new(domain.WeaponAccuracyStat)
	// 	accuracyStat.SetStat(m.Metric, m.Value)

	// 	s.log.Error("yeet", zap.Any("m", m.Metric), zap.Any("v", m.Value))

	// 	out = append(out, domain.WeaponStats{
	// 		WeaponID:      m.WeaponID,
	// 		Weapon:        m.Weapon,
	// 		Stats:         weaponStat,
	// 		AccuracyStats: accuracyStat,
	// 	})
	// }

	return weaponStats, nil
}

func (s *statistic) GetWeaponClassStats(ctx context.Context, steamID uint64, classID uint8) ([]domain.WeaponClassStats, error) {
	metrics, err := s.repo.GetWeaponClassStats(ctx, steamID, classID)
	if err != nil {
		return nil, err
	}

	if len(metrics) == 0 {
		return nil, domain.ErrWeaponClassStatsNotFound
	}

	var out []domain.WeaponClassStats

	for _, m := range metrics {
		i := sort.Search(len(out), func(i int) bool {
			return out[i].Class == m.Class
		})

		if i < len(out) && out[i].Class == m.Class {
			out[i].Stats.SetStat(m.Metric, m.Value)
			continue
		}

		s := new(domain.WeaponStat)
		s.SetStat(m.Metric, m.Value)

		out = append(out, domain.WeaponClassStats{
			ClassID: m.ClassID,
			Class:   m.Class,
			Stats:   s,
		})

	}

	return out, nil
}
