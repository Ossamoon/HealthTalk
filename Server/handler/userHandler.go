package handler


import (
    "net/http"
    "time"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    UserRespose struct {
        ID                  uint            `json:"user_id"`    
        CreatedAt           time.Time       `json:"created_at"`
        UpdatedAt           time.Time       `json:"updated_at"`
        DeletedAt           gorm.DeletedAt  `json:"deleted_at"`
        Name                string          `json:"name"`
        Friends             []UserSummary   `json:"friends"`
        ManagingGroups      []GroupSummary  `json:"managing_groups"`
        PerticipatingGroups []GroupSummary  `json:"perticipating_groups"`
    }

    UserSummary struct {
        ID          uint        `json:"user_id"`
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

    responce := UserRespose {
        ID: user.Model.ID,
        CreatedAt: user.Model.CreatedAt,
        UpdatedAt: user.Model.UpdatedAt,
        DeletedAt: user.Model.DeletedAt,
        Name: user.Name,
        Friends: friends,
    }

	return c.JSON(http.StatusOK, responce)
}