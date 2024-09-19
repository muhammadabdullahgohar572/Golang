package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlDns="root:ilove1382005#@Tcp(localhost:3306)/mysqldb";
var Database  *gorm.DB
var err error


func DataMigration()  {
	
	Database, err = gorm.Open(mysql.Open(urlDns),&gorm.Config{});
	if err!=nil{
        panic("Failed to connect to database"+error.Error(err))
    }
	Database.AutoMigrate(&Employee{})
}