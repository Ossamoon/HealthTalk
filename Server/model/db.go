package model


import (
    "fmt"
    "os"
    "time"

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

    db.AutoMigrate(&User{}, &Group{}, &DirectMessage{}, &GroupMessage{}, &HealthRecord{})

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

    // Create sample direct messages
    dm1 := DirectMessage{FromUserID: nobi.Model.ID, ToUserID: dora.Model.ID, Content: "おやつ食べれるよ"}
    dm2 := DirectMessage{FromUserID: dora.Model.ID, ToUserID: nobi.Model.ID, Content: "分かったー"}
    dm3 := DirectMessage{FromUserID: nobi.Model.ID, ToUserID: dora.Model.ID, Content: "今日はどら焼きみたい"}
    dm4 := DirectMessage{FromUserID: dora.Model.ID, ToUserID: nobi.Model.ID, Content: "やった！！"}
    dm5 := DirectMessage{FromUserID: gian.Model.ID, ToUserID: nobi.Model.ID, Content: "おいのび太、野球するぞ"}
    dm6 := DirectMessage{FromUserID: gian.Model.ID, ToUserID: nobi.Model.ID, Content: "遅れたらぶん殴るからな"}
    db.Create(&dm1)
    db.Create(&dm2)
    db.Create(&dm3)
    db.Create(&dm4)
    db.Create(&dm5)
    db.Create(&dm6)

    // Create sample group messages
    gm1 := GroupMessage{FromUserID: papa.Model.ID, ToGroupID: nobiGroup.Model.ID, Content: "今月はボーナスだぞ！"}
    gm2 := GroupMessage{FromUserID: nobi.Model.ID, ToGroupID: nobiGroup.Model.ID, Content: "やった！ゲーム買って！"}
    gm3 := GroupMessage{FromUserID: mama.Model.ID, ToGroupID: nobiGroup.Model.ID, Content: "今回は貯金です。"}
    gm4 := GroupMessage{FromUserID: dora.Model.ID, ToGroupID: nobiGroup.Model.ID, Content: "残念だったね、のび太くん"}
    db.Create(&gm1)
    db.Create(&gm2)
    db.Create(&gm3)
    db.Create(&gm4)

    // Create sample health records
    date1 := time.Date(2021, time.March, 3, 0, 0, 0, 0, time.UTC)
    date2 := time.Date(2021, time.March, 4, 0, 0, 0, 0, time.UTC)
    health1 := &HealthRecord{UserID: dora.Model.ID, Date: date1, BodyTemperature: 35.5}
    health2 := &HealthRecord{UserID: dora.Model.ID, Date: date2, BodyTemperature: 35.8, Memo: "どら焼き美味しかった"}
    health3 := &HealthRecord{UserID: nobi.Model.ID, Date: date1, BodyTemperature: 36.5, HeartRate: 86, CoughOrSoreThroat: 0, Headache: 0, Stomachache: 0, FeelTired: 0}
    health4 := &HealthRecord{UserID: nobi.Model.ID, Date: date2, BodyTemperature: 36.4, HeartRate: 84, CoughOrSoreThroat: 0, Headache: 0.15, Stomachache: 0, FeelTired: 0.45, Memo: "野球で疲れた"}
    db.Create(health1)
    db.Create(health2)
    db.Create(health3)
    db.Create(health4)
}