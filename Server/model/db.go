package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open("mysql", "../db/sample.db")
    if err != nil {
      panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}