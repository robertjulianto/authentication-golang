package repositories

import (
	"ims/database"
	"ims/types"
)

type accessRepository struct {
	db database.Database
}

type AccessRepository interface {
	CreateAccess(name string, code string) error
	GetAllAccesses() []types.Access
	GetAccessByID(ID int) *types.Access
	GetAccessByName(name string) *types.Access
	GetAccessByCode(code string) *types.Access
	DeleteAccessByID(ID int) error
}

func NewAccessRepository(db database.Database) *accessRepository {
	return &accessRepository{db}
}

func (repo *accessRepository) CreateAccess(name string, code string) error {
	access := types.Access{
		Name: name,
		Code: code,
	}
	err := repo.db.GetInstance().Create(&access).Error
	return err
}

func (repo *accessRepository) GetAllAccesses() []types.Access {
	var accesses []types.Access
	repo.db.GetInstance().Find(&accesses)
	return accesses
}

func (repo *accessRepository) GetAccessByID(ID int) *types.Access {
	var access *types.Access
	repo.db.GetInstance().Find(&access, ID)
	if access.ID == 0 {
		return nil
	}
	return access
}

func (repo *accessRepository) GetAccessByName(name string) *types.Access {
	var access *types.Access
	repo.db.GetInstance().Model(&types.Access{}).Where("name = ?", name).First(&access)
	if access.ID == 0 {
		return nil
	}
	return access
}

func (repo *accessRepository) GetAccessByCode(code string) *types.Access {
	var access *types.Access
	repo.db.GetInstance().Model(&types.Access{}).Where("code = ?", code).First(&access)
	if access.ID == 0 {
		return nil
	}
	return access
}

func (repo *accessRepository) DeleteAccessByID(ID int) error {
	err := repo.db.GetInstance().Delete(&types.Access{}, ID).Error
	return err
}
