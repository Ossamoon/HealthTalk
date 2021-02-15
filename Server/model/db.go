package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
    var err error
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "myrootpass"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME   := "mysql"

    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
      panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}