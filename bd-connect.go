package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDatabase = "root:ilove1382005#@tcp(localhost:3306)/mysqldb"

var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDatabase), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database please check Abdullah: " + err.Error())
	}
	Database.AutoMigrate(&Employee{})
}
