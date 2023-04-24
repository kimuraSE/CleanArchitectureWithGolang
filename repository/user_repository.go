package repository

import (
	"api-go/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User,email string) error
	CreateUser(user *model.User) error
}
type userRepositroy struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepositroy{db}
}

func (ur *userRepositroy)GetUserByEmail(user *model.User,email string) error {
	if err:=ur.db.Where("email = ?",email).First(user).Error;err!=nil {
		return err
	}
	return nil
}

func(ur *userRepositroy) CreateUser(user *model.User) error {
	if err:=ur.db.Create(user).Error;err!=nil {
		return err
	}
	return nil
}