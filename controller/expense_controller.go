package controller

import (
	"finances-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExpenseController interface {
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type ExpenseControllerImpl struct {
	service service.ExpenseService
}

func NewExpenseController(service service.ExpenseService) ExpenseController {
	return &ExpenseControllerImpl{service: service}
}

func (t ExpenseControllerImpl) FindById(c *gin.Context) {
	expenseId := c.Param("id")
	expense := t.service.GetById(expenseId)
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

func (t ExpenseControllerImpl) FindAll(c *gin.Context) {
	expenses := t.service.GetAll()
	c.JSON(http.StatusOK, gin.H{"data": expenses})
}
