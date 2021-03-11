package handler


import (
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


func AddHealthRecord(c echo.Context) error {
    healthRecord := new(model.HealthRecord)
    if err := c.Bind(healthRecord); err != nil {
        return err
    }

    userID := userIDFromToken(c)
    if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

    healthRecord.UserID = userID
    model.CreateHealthRecord(healthRecord)

    responce := CommonCreateResponce {
        ID: healthRecord.Model.ID,
        CreatedAt: healthRecord.Model.CreatedAt,
    }

    return c.JSON(http.StatusCreated, responce)
}


func GetHealthRecords(c echo.Context) error {
    userID := userIDFromToken(c)
    if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

    healthRecords := model.FindHealthRecords(&model.HealthRecord{UserID: userID})

    return c.JSON(http.StatusOK, healthRecords)
}