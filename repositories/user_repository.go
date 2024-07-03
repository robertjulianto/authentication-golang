package repositories

import (
	"ims/database"
	"ims/types"
)

type userRepository struct {
	db database.Database
}

type UserRepository interface {
	CreateUser(userName string, password string, displayName string) error
	GetAllUsers() []types.User
	GetUserByID(ID int) types.User
	GetUserByUserName(userName string) types.User
	UpdateUser(user *types.User) error
	DeleteUserByID(ID int) error
}

func NewUserRepository(db database.Database) *userRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(userName string, password string, displayName string) error {
	user := types.User{
		Username:    userName,
		Password:    password,
		DisplayName: displayName,
	}
	err := repo.db.GetInstance().Create(&user).Error
	return err
}

func (repo *userRepository) GetAllUsers() []types.User {
	var users []types.User
	repo.db.GetInstance().Find(&users)
	return users
}

func (repo *userRepository) GetUserByID(ID int) types.User {
	var user types.User
	repo.db.GetInstance().Find(&user, ID)
	return user
}

func (repo *userRepository) GetUserByUserName(userName string) types.User {
	var user types.User
	repo.db.GetInstance().Where("username = ?", userName).First(&user)
	return user
}

func (repo *userRepository) UpdateUser(user *types.User) error {
	err := repo.db.GetInstance().Model(&user).Updates(user).Error
	return err
}

func (repo *userRepository) DeleteUserByID(ID int) error {
	var user types.User
	err := repo.db.GetInstance().Delete(&user, ID).Error
	return err
}
