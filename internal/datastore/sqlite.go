package datastore

import (
	"time"

	"github.com/pidanou/lifetrack/internal/types"
	"gorm.io/gorm"
)

type SqliteDatastore struct {
	Db *gorm.DB
}

// Hydration

func (d *SqliteDatastore) CreateHydration(hydration *types.Hydration) error {

	result := d.Db.Create(hydration)

	return result.Error

}

func (d *SqliteDatastore) GetHydrationByDay(day time.Time) ([]types.Hydration, error) {
	var hydrations []types.Hydration

	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())

	endOfDay := startOfDay.Add(24 * time.Hour)

	if err := d.Db.Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Find(&hydrations).Error; err != nil {
		return nil, err
	}

	return hydrations, nil

}

// Sleep

func (d *SqliteDatastore) CreateSleep(sleep *types.Sleep) error {
	result := d.Db.Create(sleep)
	return result.Error
}
