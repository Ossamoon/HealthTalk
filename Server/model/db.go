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

    db.AutoMigrate(&User{}, &Group{}, &DirectMessage{}, &GroupMessage{})

    CreateSampleDataSet()
}

func CreateSampleDataSet() {
    // Create sample users
    dora := User{Name: "ドラえもん", Password: "dorayaki", Email: "dora22@gmail.com"}
    nobi := User{Name: "野比のび太", Password: "nobita", Email: "nobita@gmail.com"}
    gian := User{Name: "剛田武", Password: "kokoronotomo", Email: "gian@gmail.com"}
    sune := User{Name: "骨川スネ夫", Password: "x4RR89ed0?ef7G", Email: "honekawa-Jr@gmail.com"}
    sizu := User{Name: "源静香", Password: "shi2KaMy0010", Email: "shizuchan@gmail.com"}
    papa := User{Name: "野比のび助", Password: "nobisuke", Email: "nobisuke@gmail.com"}
    mama := User{Name: "野比玉子", Password: "NoTM884tama", Email: "tamako@gmail.com"}
    db.Create(&dora)
    db.Create(&nobi)
    db.Create(&gian)
    db.Create(&sune)
    db.Create(&sizu)
    db.Create(&papa)
    db.Create(&mama)

    // Append friends to users
    db.Model(&dora).Association("Friends").Append([]*User{&nobi, &papa, &mama})
    db.Model(&nobi).Association("Friends").Append([]*User{&gian, &sune, &sizu, &papa, &mama})
    db.Model(&gian).Association("Friends").Append([]*User{&sune, &sizu})
    db.Model(&sune).Association("Friends").Append([]*User{&sizu})
    db.Model(&papa).Association("Friends").Append([]*User{&mama})

    // Create sample groups
    nobiGroup := Group{
        Name: "野比家",
        Managers: []*User{&papa, &mama},
        Members: []*User{&nobi, &dora},
    }
    akichiGroup := Group{
        Name: "空き地に集まる会",
        Managers: []*User{&sune, &sizu},
        Members: []*User{&gian, &nobi, &dora},
    }
    giantsGroup := Group{
        Name: "野球ジャイアンズ",
        Managers: []*User{&gian},
        Members: []*User{&sune, &nobi},
    }
    db.Create(&nobiGroup)
    db.Create(&akichiGroup)
    db.Create(&giantsGroup)
}