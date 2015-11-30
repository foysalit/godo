package main

import (
	"os"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
)

func InitDB () (gorm.DB){
	dbAddress := os.Getenv("DB_ADDR")+"?parseTime=true"
	
	db, err := gorm.Open("mysql", dbAddress)
	db.LogMode(true)
	
	if err != nil {
		spew.Dump(err)
	}

	var todo Todo
	checkQuery := db.First(&todo)
	if checkQuery.Error != nil && checkQuery.Error != gorm.RecordNotFound {
		spew.Dump(checkQuery)
		db.CreateTable(todo)
	}

	return db
}