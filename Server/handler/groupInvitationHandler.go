package handler


import (
    "net/http"
	"strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
	GroupInvitationRequest struct {
		FromGroupID			uint	`json:"from_group_id"`
		ToUserID			uint	`json:"to_user_id"`
	}
)


func AddGroupInvitation(c echo.Context) error {
	userID := userIDFromToken(c)
    if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

	request := new(GroupInvitationRequest)
    if err := c.Bind(request); err != nil {
        return err
    }

	invitation := new(model.GroupInvitation)
	invitation.FromGroupID = request.FromGroupID
	invitation.ToUserID = request.ToUserID
	invitation.Status = model.UNREAD

	model.CreateGroupInvitation(invitation)

	responce := CommonCreateResponce {
        ID: invitation.Model.ID,
        CreatedAt: invitation.Model.CreatedAt,
    }

	return c.JSON(http.StatusCreated, responce)
}


func GetGroupInvitations(c echo.Context) error {
	toUserID := userIDFromToken(c)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

	groupInvitations := model.FindGroupInvitations(&model.GroupInvitation{ToUserID: toUserID, Status: model.UNREAD})

	return c.JSON(http.StatusOK, groupInvitations)
}


func UpdateGroupInvitationStatus(c echo.Context) error {
    toUserID := userIDFromToken(c)
    if toUser := model.FindUser(&model.User{Model: gorm.Model{ID: toUserID}}); toUser.ID == 0 {
        return echo.ErrNotFound
    }

    tempUint64, _ := strconv.ParseUint(c.Param("invitation_id"), 10, 64)
    invitationID := uint(tempUint64)
    invitation := model.FindGroupInvitation(&model.GroupInvitation{Model: gorm.Model{ID: invitationID}})
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
                Message: "invalid status value",
            }
    }

    model.UpdateGroupInvitation(&invitation, status)

    responce := CommonUpdateResponce {
        ID: invitation.Model.ID,
        UpdatedAt: invitation.Model.UpdatedAt,
    }

    return c.JSON(http.StatusOK, responce)
}