package main

import (
	"time"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	gorm.Model
	Name string `json:name gorm:"column:name"`
	Completed bool `json:completed sql:"default:false" gorm:"column:completed"`
	Due time.Time `json:due gorm:"column:due"`
}

type Todos []Todo

type TodoResponse struct {
	Error bool `json:error`
	Data *Todo `json:data`
}

type TodosResponse struct {
	Error bool `json:error`
	Data *Todos `json:data`
}

func (t Todo) TableName() string {
    return "todos"
}