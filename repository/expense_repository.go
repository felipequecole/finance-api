package repository

import (
	"errors"
	"finances-api/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	FindAll() []model.Expense
	FindById(expenseId uuid.UUID) (model.Expense, error)
}

type ExpenseRepositoryImpl struct {
	Db *gorm.DB
}

func NewExpenseRepositoryImpl(Db *gorm.DB) ExpenseRepository {
	return &ExpenseRepositoryImpl{Db: Db}
}

func (repository ExpenseRepositoryImpl) FindById(expenseId uuid.UUID) (model.Expense, error) {
	var expense model.Expense
	result := repository.Db.Find(&expense, expenseId)
	if result != nil {
		return expense, nil
	} else {
		return expense, errors.New("expense is not found")
	}
}

func (repository ExpenseRepositoryImpl) FindAll() []model.Expense {
	var expenses []model.Expense
	repository.Db.Find(&expenses)
	return expenses
}
