package types

import "gorm.io/gorm"

type RoleAccess struct {
	gorm.Model
	RoleID    int
	AccessID  int
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
