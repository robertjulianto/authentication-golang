package repositories

import (
	"ims/database"
	"ims/types"
)

type roleRepository struct {
	db database.Database
}

type RoleRepository interface {
	CreateRole(name string) error
	GetAllRoles() []types.Role
	GetRoleByID(ID int) types.Role
	UpdateRole(role *types.Role) error
	DeleteRoleByID(ID int) error
}

func NewRoleRepository(db database.Database) *roleRepository {
	return &roleRepository{db}
}

func (repo *roleRepository) CreateRole(name string) error {
	role := types.Role{
		Name: name,
	}
	err := repo.db.GetInstance().Create(&role).Error
	return err
}

func (repo *roleRepository) GetAllRoles() []types.Role {
	var roles []types.Role
	repo.db.GetInstance().Find(&roles)
	return roles
}

func (repo *roleRepository) GetRoleByID(ID int) types.Role {
	var role types.Role
	repo.db.GetInstance().Find(&role, ID)
	return role
}

func (repo *roleRepository) UpdateRole(role *types.Role) error {
	err := repo.db.GetInstance().Model(&role).Updates(role).Error
	return err
}

func (repo *roleRepository) DeleteRoleByID(ID int) error {
	var role types.Role
	err := repo.db.GetInstance().Delete(&role, ID).Error
	return err
}
