package types

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	Name       string
	AccessCode string
	CreatedBy  string
	UpdatedBy  string
	DeletedBy  string
}
