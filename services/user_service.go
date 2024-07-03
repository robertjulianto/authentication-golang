package services

import (
	"errors"
	"fmt"
	"ims/repositories"
	"ims/types"
)

type userService struct {
	userRepository repositories.UserRepository
}

type UserService interface {
	CreateUser(username string, password string, displayName string) error
	GetAllUsers() []types.User
	GetUserByID(ID int) types.User
	UpdateUser(spec UpdateUserSpec) error
	DeleteUserByID(ID int) error
}

type UpdateUserSpec struct {
	ID          int
	UserName    string
	Password    string
	DisplayName string
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (service *userService) CreateUser(username string, password string, displayName string) error {
	user := service.userRepository.GetUserByUserName(username)
	if user.ID != 0 {
		errorMsg := fmt.Sprintf("User with username %v already exists", username)
		return errors.New(errorMsg)
	}
	err := service.userRepository.CreateUser(username, password, displayName)
	return err
}

func (service *userService) GetAllUsers() []types.User {
	return service.userRepository.GetAllUsers()
}

func (service *userService) GetUserByID(ID int) types.User {
	return service.userRepository.GetUserByID(ID)
}

func (service *userService) UpdateUser(spec UpdateUserSpec) error {
	user := service.userRepository.GetUserByID(spec.ID)
	user.Username = spec.UserName
	user.Password = spec.Password
	user.DisplayName = spec.DisplayName
	err := service.userRepository.UpdateUser(&user)
	return err
}

func (service *userService) DeleteUserByID(ID int) error {
	err := service.userRepository.DeleteUserByID(ID)
	return err
}
