package types

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID    int
	RoleID    int
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
