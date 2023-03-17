package model

import "github.com/luyanakat/todo-restapi/common"

type TodoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

type TodoItemCreate struct {
	ID          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string       { return "todo_items" }
func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }
func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
