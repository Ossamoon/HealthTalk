package model

import (
    "fmt"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
    user := os.Getenv("MYSQL_USER")
    pass := os.Getenv("MYSQL_PASSWORD")
    // host := "tcp(127.0.0.1:3306)"
    host := os.Getenv("MYSQL_HOST")
    dbname := os.Getenv("MYSQL_DATABASE")
    connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

    db, err := gorm.Open("mysql", connection)
    defer db.Close()

    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}