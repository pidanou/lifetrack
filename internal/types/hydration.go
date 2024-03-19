package types

import (
	"gorm.io/gorm"
)

type Hydration struct {
	gorm.Model
	Quantity float64
	Name     string
}
