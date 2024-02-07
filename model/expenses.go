package model

import "github.com/google/uuid"

type Expense struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key"`
	Name string    `gorm:"type:varchar(255)"`
}
