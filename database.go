package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySQLDBConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s/@%s/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBName))
	return
}
