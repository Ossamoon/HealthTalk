package handler


import (
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    UserRespose struct {
        gorm.Model
        Name                string          `json:"name"`
        Friends             []UserSummary   `json:"friends"`
        ManagingGroups      []GroupSummary  `json:"managing_groups"`
        PerticipatingGroups []GroupSummary  `json:"perticipating_groups"`
    }

    UserSummary struct {
        ID          uint 
        Name        string      `json:"name"`
    }
)


func GetUser(c echo.Context) error {
	userID := userIDFromToken(c)
	user := model.FindUserWithPreload(&model.User{Model: gorm.Model{ID: userID}})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    var friends []UserSummary
    for _, friend := range user.Friends {
        adding := UserSummary {
            ID: friend.Model.ID,
            Name: friend.Name,
        }
        friends = append(friends, adding)
    }

    var managingGroups []GroupSummary
    fmt.Println(user.ManagingGroups)
    for _, group := range user.ManagingGroups {
        adding := GroupSummary {
            ID: group.Model.ID,
            Name: group.Name,
        }
        managingGroups = append(managingGroups, adding)
    }

    var perticipatingGroups []GroupSummary
    fmt.Println(user.PerticipatingGroups)
    for _, group := range user.PerticipatingGroups {
        adding := GroupSummary {
            ID: group.Model.ID,
            Name: group.Name,
        }
        perticipatingGroups = append(perticipatingGroups, adding)
    }

    responce := UserRespose {
        Model: user.Model,
        Name: user.Name,
        Friends: friends,
        ManagingGroups: managingGroups,
        PerticipatingGroups: perticipatingGroups,
    }

	return c.JSON(http.StatusOK, responce)
}