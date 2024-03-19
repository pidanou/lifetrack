package datastore

import (
	"time"

	"github.com/pidanou/lifetrack/internal/types"
)

type Datastore interface {
	CreateHydration(*types.Hydration) error
	GetHydrationByDay(time.Time) ([]types.Hydration, error)

	CreateSleep(*types.Sleep) error
}
