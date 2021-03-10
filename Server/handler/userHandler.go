package handler

import (
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)

func GetUser(c echo.Context) error {
	userID := userIDFromToken(c)
	user := model.FindUserWithPreload(&model.User{Model: gorm.Model{ID: userID}})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    user.Password = ""
    for _, friend := range user.Friends {
        friend.Password = ""
        friend.Email = ""
    }

	return c.JSON(http.StatusOK, user)
}