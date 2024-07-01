package types

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name      string
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
