package repository

import (
	model "GoTutorial/Model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	AddUser(Name string, Surname string, Address string, Age int, Password string) (bool, error)
	DeleteUser(ID uint64) (bool, error)
	UpdateUserPassword(ID uint64, Password string) (bool, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {

	return &UserRepository{
		DB: DB,
	}
}

func (userRepository *UserRepository) AddUser(Name string, Surname string, Address string, Age int, Password string) (bool, error) {

	user := model.NewUser(Name, Surname, Address, Age, Password)
	result := userRepository.DB.Create(user)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (userRepository *UserRepository) DeleteUser(ID uint64) (bool, error) {

	result := userRepository.DB.Delete(&model.User{}, ID)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (userRepository *UserRepository) UpdateUserPassword(ID uint64, Password string) (bool, error) {

	var user model.User
	result := userRepository.DB.First(&user, ID)
	if result.Error != nil {
		return false, result.Error
	}

	user.Password = Password
	updateResult := userRepository.DB.Save(&user)
	if updateResult.Error != nil {
		return false, updateResult.Error
	}

	return true, nil
}
