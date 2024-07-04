package services

import (
	"ims/repositories"
	"ims/types"
)

type userRoleService struct {
	userRoleRepository repositories.UserRoleRepository
	userRepository     repositories.UserRepository
}

type UserRoleService interface {
	CreateUserRole(roleID int, userIDs []int) error
	DeleteUserRole(roleID int, userIDs []int) error
	GetRoleMembers(roleID int) []types.User
}

func NewUserRoleService(
	userRoleRepository repositories.UserRoleRepository,
	userRepository repositories.UserRepository) *userRoleService {
	return &userRoleService{userRoleRepository, userRepository}
}

func (service *userRoleService) CreateUserRole(roleID int, userIDs []int) error {
	err := service.userRoleRepository.CreateUserRole(roleID, userIDs)
	return err
}

func (service *userRoleService) DeleteUserRole(roleID int, userIDs []int) error {
	err := service.userRoleRepository.DeleteUserRole(roleID, userIDs)
	return err
}

func (service *userRoleService) GetRoleMembers(roleID int) []types.User {
	userIDs := service.userRoleRepository.GetUserIDByRoleID(roleID)
	users := service.userRepository.GetUsersByUserIDs(userIDs)
	return users
}
