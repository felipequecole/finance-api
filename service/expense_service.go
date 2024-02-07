package service

import (
	"finances-api/model"
	"finances-api/repository"
	"github.com/google/uuid"
)

type ExpenseService interface {
	GetById(expenseId string) model.Expense
	GetAll() []model.Expense
}

type ExpenseServiceImpl struct {
	repository repository.ExpenseRepository
}

func NewExpenseService(repository repository.ExpenseRepository) ExpenseService {
	return &ExpenseServiceImpl{repository: repository}
}

func (t ExpenseServiceImpl) GetById(expenseId string) model.Expense {
	dbId, _ := uuid.Parse(expenseId)
	expense, _ := t.repository.FindById(dbId)
	return expense
}

func (t ExpenseServiceImpl) GetAll() []model.Expense {
	return t.repository.FindAll()
}
