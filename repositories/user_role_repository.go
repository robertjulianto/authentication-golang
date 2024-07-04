package repositories

import (
	"ims/database"
	"ims/types"
)

type userRoleRepository struct {
	db database.Database
}

type UserRoleRepository interface {
	CreateUserRole(roleID int, userIDs []int) error
	DeleteUserRole(roleID int, userIDs []int) error
	GetUserIDByRoleID(roleID int) []int
}

func NewUserRoleRepository(db database.Database) *userRoleRepository {
	return &userRoleRepository{db}
}

func (repo *userRoleRepository) CreateUserRole(roleID int, userIDs []int) error {
	userRoles := createUserRoleObject(roleID, userIDs)
	err := repo.db.GetInstance().Create(&userRoles).Error
	return err
}

func (repo *userRoleRepository) DeleteUserRole(roleID int, userIDs []int) error {
	userRoles := createUserRoleObject(roleID, userIDs)
	userRoleIDs := []int{}
	for _, userRole := range userRoles {
		userRoleIDs = append(userRoleIDs, userRole.UserID)
	}
	err := repo.db.GetInstance().Unscoped().Where("user_id in (?)", userRoleIDs).Delete(&userRoles).Error
	return err
}

func (repo *userRoleRepository) GetUserIDByRoleID(roleID int) []int {
	var userRoles []types.UserRole
	repo.db.GetInstance().Where("role_id = ?", roleID).Find(&userRoles)
	userIDs := []int{}
	for _, userRole := range userRoles {
		userIDs = append(userIDs, userRole.UserID)
	}
	return userIDs
}

func createUserRoleObject(roleID int, userIDs []int) []types.UserRole {
	var userRole types.UserRole
	userRoles := []types.UserRole{}

	for _, userID := range userIDs {
		userRole = types.UserRole{
			UserID: userID,
			RoleID: roleID,
		}
		userRoles = append(userRoles, userRole)
	}
	return userRoles
}
