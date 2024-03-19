package types

import "gorm.io/gorm"

type Sleep struct {
	gorm.Model
	Quantity float64
}
