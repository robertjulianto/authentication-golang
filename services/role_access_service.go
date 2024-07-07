package services

import (
	"ims/repositories"
	"ims/types"
)

type roleAccessService struct {
	roleAccessRepository repositories.RoleAccessRepository
	accessRepository     repositories.AccessRepository
}

type RoleAccessService interface {
	CreateRoleAccess(roleID int, userIDs []int) error
	DeleteRoleAccess(roleID int, userIDs []int) error
	GetAccessByRoleID(roleID int) []types.Access
}

func NewRoleAccessService(
	roleAccessRepository repositories.RoleAccessRepository,
	accessRepository repositories.AccessRepository) *roleAccessService {
	return &roleAccessService{roleAccessRepository, accessRepository}
}

func (service *roleAccessService) CreateRoleAccess(roleID int, accessIDs []int) error {
	err := service.roleAccessRepository.CreateRoleAccess(roleID, accessIDs)
	return err
}

func (service *roleAccessService) DeleteRoleAccess(roleID int, accessIDs []int) error {
	err := service.roleAccessRepository.DeleteRoleAccess(roleID, accessIDs)
	return err
}

func (service *roleAccessService) GetAccessByRoleID(roleID int) []types.Access {
	accessIDs := service.roleAccessRepository.GetAccessIDByRoleID(roleID)
	accesses := service.accessRepository.GetAccessesByAccessIDs(accessIDs)
	return accesses
}
