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
    host := os.Getenv("MYSQL_HOST")
    dbname := os.Getenv("MYSQL_DATABASE")
    connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

    var err error
    db, err = gorm.Open("mysql", connection)
    fmt.Println("Set db:", db)

    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}