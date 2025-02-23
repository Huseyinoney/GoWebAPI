package service

import (
	repository "GoTutorial/Repository"
	"fmt"
)

type IUserService interface {
	AddUser(Name string, Surname string, Address string, Age int, Password string) bool
	DeleteUser(ID uint64) bool
	UpdateUserPassword(ID uint64, Password string) bool
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {

	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) AddUser(Name string, Surname string, Address string, Age int, Password string) bool {

	result, error := userService.userRepository.AddUser(Name, Surname, Address, Age, Password)
	if !result {
		fmt.Println(error)
		return false
	}
	return true
}

func (userService *UserService) DeleteUser(ID uint64) bool {

	result, error := userService.userRepository.DeleteUser(ID)
	if !result {
		fmt.Println(error)
		return false
	}
	return true
}

func (userService *UserService) UpdateUserPassword(ID uint64, Password string) bool {

	result, error := userService.userRepository.UpdateUserPassword(ID, Password)
	if !result {
		fmt.Println(error)
		return false
	}
	return true
}
