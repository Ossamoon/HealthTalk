package handler


import (
    "net/http"
	"strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    GroupResponce struct {
        gorm.Model
        Name        string          `json:"name"`
        Managers    []UserSummary   `json:"manages"`
        Members     []UserSummary   `json:"members"`
    }

    GroupSummary struct {
        ID          uint
        Name        string      `json:"name"`
    }
)


func AddGroup(c echo.Context) error {
    group := new(model.Group)
    if err := c.Bind(group); err != nil {
        return err
    }

    managerID := userIDFromToken(c)
    manager := model.FindUser(&model.User{Model: gorm.Model{ID: managerID}})
    if manager.ID == 0 {
        return echo.ErrNotFound
    }

    group.Managers = []*model.User{&manager}
    model.CreateGroup(group)

    responce := CommonCreateResponce {
        ID: group.Model.ID,
        CreatedAt: group.Model.CreatedAt,
    }

    return c.JSON(http.StatusCreated, responce)
}


func GetGroup(c echo.Context) error {
	tempUint64, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
    groupID := uint(tempUint64)
	group := model.FindGroupWithPreload(&model.Group{Model: gorm.Model{ID: groupID}})
    if group.ID == 0 {
        return echo.ErrNotFound
    }

    var managers []UserSummary
    for _, manager := range group.Managers {
        adding := UserSummary {
            ID: manager.Model.ID,
            Name: manager.Name,
        }
        managers = append(managers, adding)
    }
    
    var members []UserSummary
    for _, member := range group.Members {
        adding := UserSummary {
            ID: member.Model.ID,
            Name: member.Name,
        }
        members = append(members, adding)
    }

    responce := GroupResponce {
        Model: group.Model,
        Name: group.Name,
        Managers: managers,
        Members: members,
    }

	return c.JSON(http.StatusOK, responce)
}