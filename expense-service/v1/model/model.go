package model

import (
	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
	"time"
)

type Category struct {
	rdbms.Entity

	ID       uint32 `gorm:"primary_key"`
	Limit    float32
	Name     string
	Expenses []Expense
	UserID   uint32
}

type Expense struct {
	rdbms.Entity

	ID         uint32 `gorm:"primary_key"`
	Date       time.Time
	Value      float32
	Category   Category
	CategoryID uint32
	UserID     uint32
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

func (cr *CategoryRepository) FindByUserID(userID uint32) ([]Category, error) {
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

func (er *ExpenseRepository) FindByID(id uint32) (*Expense, error) {
	var expense Expense

	context := er.Context().(*gorm.DbContext)

	if err := context.DB.Preload("Category").First(&expense, Expense{ID: id}).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}
