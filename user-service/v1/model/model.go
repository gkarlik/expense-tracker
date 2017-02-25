package model

import (
	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
)

type User struct {
	rdbms.Entity

	ID        uint   `gorm:"primary_key"`
	Login     string `gorm:"type:varchar(100);unique_index"`
	GUID      string `gorm:"type:varchar(100);unique_index"`
	Password  string `gorm:"size:100"`
	Pin       string `gorm:"size:100"`
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:200"`
}

type UserRepository struct {
	*gorm.RepositoryBase
}

func (ur *UserRepository) FindByGUID(guid string) (*User, error) {
	var user User
	if err := ur.First(&user, User{GUID: guid}); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByLogin(login string) (*User, error) {
	var user User
	if err := ur.First(&user, User{Login: login}); err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(c rdbms.DbContext) *UserRepository {
	repo := &UserRepository{
		RepositoryBase: &gorm.RepositoryBase{},
	}
	repo.SetContext(c)

	return repo
}
