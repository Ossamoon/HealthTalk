package handler

import (
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)

func GetUser(c echo.Context) error {
    userID := userIDFromToken(c)
	user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    return c.JSON(http.StatusOK, user)
}

// func GetFriends(c echo.Context) error {
// 	userID := userIDFromToken(c)
// 	user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}})
//     if user.ID == 0 {
//         return echo.ErrNotFound
//     }

// 	friends := []User
// 	db.Model(&user).Association("Friends").Find(&friends)

// 	return c.JSON(http.StatusOK, friends)
// }