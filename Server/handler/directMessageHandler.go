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

    fromUserID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUserID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

	toUserID := directMessage.ToUserID
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessage.FromUserID = fromUserID
    model.CreateDirectMessage(directMessage)

    return c.JSON(http.StatusCreated, directMessage)
}

func GetDirectMessages(c echo.Context) error {
    fromUserID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUserID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessages := model.FindDirectMessages(&model.DirectMessage{FromUserID: fromUserID})
    return c.JSON(http.StatusOK, directMessages)
}