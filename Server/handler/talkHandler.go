package handler

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/Ossamoon/HealthTalk/Server/model"
)

func AddTalk(c echo.Context) error {
    talk := new(model.Talk)
    if err := c.Bind(talk); err != nil {
        return err
    }

    if talk.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid to or message fields",
        }
    }

    uid_from := userIDFromToken(c)
    if user_from := model.FindUser(&model.User{ID: uid_from}); user_from.ID == 0 {
        return echo.ErrNotFound
    }

	uid_to := talk.UID_to
    if user_to := model.FindUser(&model.User{ID: uid_to}); user_to.ID == 0 {
        return echo.ErrNotFound
    }

    talk.UID_from = uid_from
    model.CreateTalk(talk)

    return c.JSON(http.StatusCreated, talk)
}

func GetTalks(c echo.Context) error {
    uid_from := userIDFromToken(c)
    if user_from := model.FindUser(&model.User{ID: uid_from}); user_from.ID == 0 {
        return echo.ErrNotFound
    }

    talks := model.FindTalks(&model.Talk{UID_from: uid_from})
    return c.JSON(http.StatusOK, talks)
}