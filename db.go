package main

import (
	"os"
	"bytes"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
)

func InitDB () (gorm.DB){
	var dbStringBuffer bytes.Buffer
	dbStringBuffer.WriteString(os.Getenv("DB_USER"))
	dbStringBuffer.WriteString(":")
	dbStringBuffer.WriteString(os.Getenv("DB_PASS"))
	dbStringBuffer.WriteString("@/")
	dbStringBuffer.WriteString(os.Getenv("DB_NAME"))
	dbStringBuffer.WriteString("?parseTime=true")
	
	db, err := gorm.Open("mysql", dbStringBuffer.String())
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