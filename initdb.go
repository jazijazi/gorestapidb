package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db * gorm.DB
var err error
const dns = "host=localhost user=jazisql password=123456789 dbname=gormdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func InitialDatabase(){

	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("Cannot Connect to DB")
	}
}