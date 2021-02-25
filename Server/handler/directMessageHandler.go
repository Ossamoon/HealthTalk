package handler

import (
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)

func AddDirectMessage(c echo.Context) error {
    directMessage := new(model.DirectMessage)
    if err := c.Bind(directMessage); err != nil {
        return err
    }

    if directMessage.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid to or message fields",
        }
    }

    fromUID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

	toUID := directMessage.ToUID
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessage.FromUID = fromUID
    model.CreateDirectMessage(directMessage)

    return c.JSON(http.StatusCreated, directMessage)
}

func GetDirectMessages(c echo.Context) error {
    fromUID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessages := model.FindDirectMessages(&model.DirectMessage{FromUID: fromUID})
    return c.JSON(http.StatusOK, directMessages)
}