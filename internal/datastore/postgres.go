package datastore

import (
	"fmt"
	"os"
	"time"

	"github.com/pidanou/lifetrack/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresDatastore struct {
	Db *gorm.DB
}

func NewPostgresDatastore() *PostgresDatastore {

	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var dbname = os.Getenv("DB_NAME")
	var dbport = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("failed to connect database")
	}

	return &PostgresDatastore{Db: db}
}

// Hydration

func (d *PostgresDatastore) CreateHydration(hydration *types.Hydration) error {

	result := d.Db.Create(hydration)

	return result.Error

}

func (d *PostgresDatastore) GetHydrationByDay(day time.Time) ([]types.Hydration, error) {
	var hydrations []types.Hydration

	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())

	endOfDay := startOfDay.Add(24 * time.Hour)

	if err := d.Db.Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Find(&hydrations).Error; err != nil {
		return nil, err
	}

	return hydrations, nil

}

// Sleep

func (d *PostgresDatastore) CreateSleep(sleep *types.Sleep) error {
	result := d.Db.Create(sleep)
	return result.Error
}
