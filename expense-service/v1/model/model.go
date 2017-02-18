package model

import (
	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
	"time"
)

type Category struct {
	rdbms.Entity

	ID       uint `gorm:"primary_key"`
	Limit    float64
	Name     string
	Expenses []Expense
	UserID   uint
}

type Expense struct {
	rdbms.Entity

	ID         uint `gorm:"primary_key"`
	Date       time.Time
	Value      float64
	Category   Category
	CategoryID uint
	UserID     uint
}

type CategoryRepository struct {
	*gorm.RepositoryBase
}

func NewCategoryRepository(c rdbms.DbContext) *CategoryRepository {
	repo := &CategoryRepository{
		RepositoryBase: &gorm.RepositoryBase{},
	}
	repo.SetContext(c)

	return repo
}

func (cr *CategoryRepository) FindByUserID(userID uint) ([]Category, error) {
	var categories []Category
	if err := cr.Find(&categories, Category{UserID: userID}); err != nil {
		return nil, err
	}
	return categories, nil
}

type ExpenseRepository struct {
	*gorm.RepositoryBase
}

func NewExpenseRepository(c rdbms.DbContext) *ExpenseRepository {
	repo := &ExpenseRepository{
		RepositoryBase: &gorm.RepositoryBase{},
	}
	repo.SetContext(c)

	return repo
}

func (er *ExpenseRepository) FindByUserID(userID uint) ([]Expense, error) {
	var expenses []Expense
	if err := er.Find(&expenses, Expense{UserID: userID}); err != nil {
		return nil, err
	}
	return expenses, nil
}
