package handler


import (
    "net/http"
	"strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


func AddFriendInvitation(c echo.Context) error {
	fromUserID := userIDFromToken(c)
    if fromUser := model.FindUser(&model.User{Model: gorm.Model{ID: fromUserID}}); fromUser.ID == 0 {
        return echo.ErrNotFound
    }

	tempUint64, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)
    toUserID := uint(tempUint64)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

	invitation := new(model.FriendInvitation)
	invitation.FromUserID = fromUserID
	invitation.ToUserID = toUserID
	invitation.Status = model.UNREAD

	model.CreateFriendInvitation(invitation)

	responce := CommonCreateResponce {
        ID: invitation.Model.ID,
        CreatedAt: invitation.Model.CreatedAt,
    }

	return c.JSON(http.StatusCreated, responce)
}


func GetFriendInvitations(c echo.Context) error {
	toUserID := userIDFromToken(c)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

	friendInvitations := model.FindFriendInvitations(&model.FriendInvitation{ToUserID: toUserID, Status: model.UNREAD})

	return c.JSON(http.StatusCreated, friendInvitations)
}
