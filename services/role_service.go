package services

import (
	"ims/repositories"
	"ims/types"
)

type roleService struct {
	roleRepository repositories.RoleRepository
}

type RoleService interface {
	CreateRole(name string) error
	GetAllRoles() []types.Role
	GetRoleByID(ID int) types.Role
	UpdateRole(spec UpdateRoleSpec) error
	DeleteRoleByID(ID int) error
}

type UpdateRoleSpec struct {
	ID       int
	RoleName string
}

func NewRoleService(roleRepository repositories.RoleRepository) *roleService {
	return &roleService{roleRepository}
}

func (service *roleService) CreateRole(name string) error {
	err := service.roleRepository.CreateRole(name)
	return err
}

func (service *roleService) GetAllRoles() []types.Role {
	return service.roleRepository.GetAllRoles()
}

func (service *roleService) GetRoleByID(ID int) types.Role {
	return service.roleRepository.GetRoleByID(ID)
}

func (service *roleService) UpdateRole(spec UpdateRoleSpec) error {
	role := service.roleRepository.GetRoleByID(spec.ID)
	role.Name = spec.RoleName
	err := service.roleRepository.UpdateRole(&role)
	return err
}

func (service *roleService) DeleteRoleByID(ID int) error {
	err := service.roleRepository.DeleteRoleByID(ID)
	return err
}
