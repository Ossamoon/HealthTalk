package model

import (
    "fmt"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var db *gorm.DB

func init() {
    user := os.Getenv("MYSQL_USER")
    pass := os.Getenv("MYSQL_PASSWORD")
    host := os.Getenv("MYSQL_HOST")
    dbname := os.Getenv("MYSQL_DATABASE")
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbname)

    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    fmt.Println("Set db:", db)

    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &Group{}, &DirectMessage{})

    CreateSampleDataSet()
}

func CreateSampleDataSet() {
    // Create sample users
    db.Create(&User{Name: "ドラえもん", Password: "dorayaki", Email: "dora22@gmail.com"})
    db.Create(&User{Name: "野比のび太", Password: "nobita", Email: "nobita@gmail.com"})
    db.Create(&User{Name: "剛田武", Password: "kokoronotomo", Email: "gian@gmail.com"})
    db.Create(&User{Name: "骨川スネ夫", Password: "x4RR89ed0?ef7G", Email: "honekawa-Jr@gmail.com"})
    db.Create(&User{Name: "源静香", Password: "shi2KaMy0010", Email: "shizuchan@gmail.com"})
    db.Create(&User{Name: "野比のび助", Password: "nobisuke", Email: "nobisuke@gmail.com"})
    db.Create(&User{Name: "野比玉子", Password: "NoTM884tama", Email: "tamako@gmail.com"})

    //Create sample groups
    db.Create(&Group{Name: "野比家"})
    db.Create(&Group{Name: "空き地に集まる会"})
    db.Create(&Group{Name: "剛田様 特別ご招待"})
}