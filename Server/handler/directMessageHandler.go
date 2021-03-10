package handler


import (
    "net/http"
    "strconv"
    "sort"

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

	tempUint64, _ := strconv.ParseUint(c.Param("with"), 10, 64)
    toUserID := uint(tempUint64)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessage.FromUserID = fromUserID
    directMessage.ToUserID = toUserID
    model.CreateDirectMessage(directMessage)

    return c.JSON(http.StatusCreated, directMessage)
}


func GetDirectMessages(c echo.Context) error {
    fromUserID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUserID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

    tempUint64, _ := strconv.ParseUint(c.Param("with"), 10, 64)
    toUserID := uint(tempUint64)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    directMessages := model.FindDirectMessages(&model.DirectMessage{FromUserID: fromUserID, ToUserID: toUserID})
    directMessages = append(directMessages, model.FindDirectMessages(&model.DirectMessage{FromUserID: toUserID, ToUserID: fromUserID})...)
    sort.Slice(directMessages, func(i, j int) bool {
        return directMessages[i].Model.ID < directMessages[j].Model.ID
    })

    return c.JSON(http.StatusOK, directMessages)
}