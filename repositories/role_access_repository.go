package repositories

import (
	"ims/database"
	"ims/types"
)

type roleAccessRepository struct {
	db database.Database
}

type RoleAccessRepository interface {
	CreateRoleAccess(roleID int, userIDs []int) error
	DeleteRoleAccess(roleID int, userIDs []int) error
	GetAccessIDByRoleID(roleID int) []int
}

func NewRoleAccessRepository(db database.Database) *roleAccessRepository {
	return &roleAccessRepository{db}
}

func (repo *roleAccessRepository) CreateRoleAccess(roleID int, userIDs []int) error {
	roleAccesses := createRoleAccessObject(roleID, userIDs)
	err := repo.db.GetInstance().Create(&roleAccesses).Error
	return err
}

func (repo *roleAccessRepository) DeleteRoleAccess(roleID int, accessIDs []int) error {
	err := repo.db.GetInstance().Unscoped().Where("role_id = (?) AND access_id IN (?)", roleID, accessIDs).Delete(&types.RoleAccess{}).Error
	return err
}

func (repo *roleAccessRepository) GetAccessIDByRoleID(roleID int) []int {
	var roleAccesses []types.RoleAccess
	repo.db.GetInstance().Where("role_id = ?", roleID).Find(&roleAccesses)
	userIDs := []int{}
	for _, roleAccess := range roleAccesses {
		userIDs = append(userIDs, roleAccess.AccessID)
	}
	return userIDs
}

func createRoleAccessObject(roleID int, accessIDs []int) []types.RoleAccess {
	var roleAccess types.RoleAccess
	roleAccesses := []types.RoleAccess{}

	for _, accessID := range accessIDs {
		roleAccess = types.RoleAccess{
			RoleID:   roleID,
			AccessID: accessID,
		}
		roleAccesses = append(roleAccesses, roleAccess)
	}
	return roleAccesses
}
