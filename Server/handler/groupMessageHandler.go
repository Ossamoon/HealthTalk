package handler


import (
    "net/http"
    "strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


func AddGroupMessage(c echo.Context) error {
    groupMessage := new(model.GroupMessage)
    if err := c.Bind(groupMessage); err != nil {
        return err
    }

    if groupMessage.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid to or message fields",
        }
    }

    fromUserID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUserID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

	tempUint64, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
    toGroupID := uint(tempUint64)
    if toGroup := model.FindGroup(&model.Group{Model: gorm.Model{ID: toGroupID}}); toGroup.ID == 0 {
        return echo.ErrNotFound
    }

    groupMessage.FromUserID = fromUserID
    groupMessage.ToGroupID = toGroupID
    model.CreateGroupMessage(groupMessage)

    return c.JSON(http.StatusCreated, groupMessage)
}


func GetGroupMessages(c echo.Context) error {
    tempUint64, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
    toGroupID := uint(tempUint64)
    if toGroup := model.FindGroup(&model.Group{Model: gorm.Model{ID: toGroupID}}); toGroup.ID == 0 {
        return echo.ErrNotFound
    }

    groupMessages := model.FindGroupMessages(&model.GroupMessage{ToGroupID: toGroupID})

    return c.JSON(http.StatusOK, groupMessages)
}