package handler


import (
    "time"
    "net/http"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    UserRespose struct {
        gorm.Model
        Name                string          `json:"name"`
        Email               string          `json:"email"`
        Friends             []UserSummary   `json:"friends"`
        ManagingGroups      []GroupSummary  `json:"managing_groups"`
        PerticipatingGroups []GroupSummary  `json:"perticipating_groups"`
    }

    UpdateUserRequest struct {
        Name        string      `json:"name"`
        Email       string      `json:"email"`
    }

    UpdateUserResponse struct {
        ID          uint
        UpdatedAt   time.Time
        Name        string      `json:"name"`
        Email       string      `json:"email"`
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
    for _, group := range user.ManagingGroups {
        adding := GroupSummary {
            ID: group.Model.ID,
            Name: group.Name,
        }
        managingGroups = append(managingGroups, adding)
    }

    var perticipatingGroups []GroupSummary
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
        Email: user.Email,
        Friends: friends,
        ManagingGroups: managingGroups,
        PerticipatingGroups: perticipatingGroups,
    }

	return c.JSON(http.StatusOK, responce)
}

func UpdateUser(c echo.Context) error {
    userID := userIDFromToken(c)
	user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    updating := new(UpdateUserRequest)
    if err := c.Bind(updating); err != nil {
        return err
    }

    model.UpdateUser(&user, updating.Name, updating.Email)

    responce := UpdateUserResponse {
        ID: user.Model.ID,
        UpdatedAt: user.Model.UpdatedAt,
        Name: user.Name,
        Email: user.Email,
    }

    return c.JSON(http.StatusOK, responce)
}