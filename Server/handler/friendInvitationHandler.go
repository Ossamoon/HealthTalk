package handler


import (
    "net/http"
	"strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    UpdateInvitationStatusRequest struct {
        Status      string      `json:"status"`
    }
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


func UpdateFriendInvitationStatus(c echo.Context) error {
    toUserID := userIDFromToken(c)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    tempUint64, _ := strconv.ParseUint(c.Param("invitation_id"), 10, 64)
    invitationID := uint(tempUint64)
    invitation := model.FindFriendInvitation(&model.FriendInvitation{Model: gorm.Model{ID: invitationID}})
    if invitation.ID == 0 {
        return echo.ErrNotFound
    }
    if invitation.ToUserID != toUserID {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "this invitation is not for you",
        }
    }
    if invitation.Status == model.ACCEPTED || invitation.Status == model.REFUSED {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "this invitation has already been accepted or refused",
        }
    }

    updating := new(UpdateInvitationStatusRequest)
    if err := c.Bind(updating); err != nil {
        return err
    }

    var status uint8
    switch updating.Status {
        case "pending":
            status = model.PENDING
        case "accepted":
            status = model.ACCEPTED
        case "refused":
            status = model.REFUSED
        default:
            return &echo.HTTPError{
                Code:    http.StatusBadRequest,
                Message: "this invitation is not for you",
            }
    }

    model.UpdateFriendInvitation(&invitation, status)

    responce := CommonUpdateResponce {
        ID: invitation.Model.ID,
        UpdatedAt: invitation.Model.UpdatedAt,
    }

    return c.JSON(http.StatusOK, responce)
}
