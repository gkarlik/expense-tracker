package model

import (
	"time"

	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
)

type Category struct {
	rdbms.Entity

	ID       string `gorm:"primary_key;type:varchar(100);"`
	Limit    float32
	Name     string
	Expenses []Expense
	UserID   string `gorm:"type:varchar(100)"`
}

func (c Category) IsValid() bool {
	if c.ID == "" || c.Name == "" || c.UserID == "" || c.Limit <= 0 {
		return false
	}
	return true
}

type Expense struct {
	rdbms.Entity

	ID         string `gorm:"primary_key;type:varchar(100)"`
	Date       time.Time
	Value      float32
	Category   Category
	CategoryID string `gorm:"type:varchar(100)"`
	UserID     string `gorm:"type:varchar(100)"`
}

func (e Expense) IsValid() bool {
	if e.ID == "" || e.CategoryID == "" || e.UserID == "" || e.Value <= 0 || e.Date.Unix() > time.Now().Unix() || e.Date.Unix() < time.Date(1980, time.January, 1, 23, 0, 0, 0, time.UTC).Unix() {
		return false
	}
	return true
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

func (cr *CategoryRepository) FindByID(id string) (*Category, error) {
	var category Category
	if err := cr.Find(&category, Category{ID: id}); err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *CategoryRepository) FindByUserID(userID string, offset, limit int32) ([]Category, error) {
	var categories []Category

	context := cr.Context().(*gorm.DbContext)
	if err := context.DB.Limit(limit).Offset(offset).Find(&categories, Category{UserID: userID}).Error; err != nil {
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

func (er *ExpenseRepository) FindByID(id string) (*Expense, error) {
	var expense Expense

	context := er.Context().(*gorm.DbContext)

	if err := context.DB.Preload("Category").First(&expense, Expense{ID: id}).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (er *ExpenseRepository) FindByUserID(userID string, offset, limit int32) ([]Expense, error) {
	var expenses []Expense

	context := er.Context().(*gorm.DbContext)
	if err := context.DB.Preload("Category").Limit(limit).Offset(offset).Find(&expenses, Expense{UserID: userID}).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}
