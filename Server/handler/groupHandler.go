package handler

import (
    "net/http"
	"strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)

func GetGroup(c echo.Context) error {
	tempUint64, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
    groupID := uint(tempUint64)
	group := model.FindGroupWithPreload(&model.Group{Model: gorm.Model{ID: groupID}})
    if group.ID == 0 {
        return echo.ErrNotFound
    }

    for _, manager := range group.Managers {
        manager.Password = ""
        manager.Email = ""
    }

	for _, member := range group.Members {
        member.Password = ""
        member.Email = ""
    }

	return c.JSON(http.StatusOK, group)
}