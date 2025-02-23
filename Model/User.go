package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func NewUser(Name string, Surname string, Address string, Age int, Password string) *User {
	return &User{
		Name:     Name,
		Surname:  Surname,
		Address:  Address,
		Age:      Age,
		Password: Password,
	}
}
