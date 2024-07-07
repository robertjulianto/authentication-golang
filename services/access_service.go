package services

import (
	"ims/repositories"
	"ims/types"
)

type accessService struct {
	accessRepository repositories.AccessRepository
}

type AccessService interface {
	CreateAccess(name string, code string) error
	GetAllAccesses() []types.Access
	GetAccessByID(ID int) *types.Access
	GetAccessByName(name string) *types.Access
	GetAccessByCode(code string) *types.Access
	DeleteAccessByID(ID int) error
}

func NewAccessService(accessRepository repositories.AccessRepository) *accessService {
	return &accessService{accessRepository}
}

func (service *accessService) CreateAccess(name string, code string) error {
	err := service.accessRepository.CreateAccess(name, code)
	return err
}

func (service *accessService) GetAllAccesses() []types.Access {
	accesses := service.accessRepository.GetAllAccesses()
	return accesses
}

func (service *accessService) GetAccessByID(ID int) *types.Access {
	access := service.accessRepository.GetAccessByID(ID)
	return access
}

func (service *accessService) GetAccessByName(name string) *types.Access {
	access := service.accessRepository.GetAccessByName(name)
	return access
}

func (service *accessService) GetAccessByCode(code string) *types.Access {
	access := service.accessRepository.GetAccessByCode(code)
	return access
}

func (service *accessService) DeleteAccessByID(ID int) error {
	err := service.accessRepository.DeleteAccessByID(ID)
	return err
}
