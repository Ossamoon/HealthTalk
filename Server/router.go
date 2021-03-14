package main


import (
	"github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
	"github.com/Ossamoon/HealthTalk/Server/handler"
)


func newRouter() *echo.Echo {
	e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/test", handler.Test)
    e.POST("/signup", handler.Signup) // POST /signup
    e.POST("/login", handler.Login) // POST /login

    api := e.Group("/api")
    api.Use(middleware.JWTWithConfig(handler.Config)) // /apiより下はJWTの認証が必要

    api.GET("/user", handler.GetUser) // GET /api/user
    api.PUT("/user", handler.UpdateUser) // PUT /api/user

    api.GET("/group/:group_id", handler.GetGroup) // GET /api/group/:group_id
    api.POST("/group", handler.AddGroup) // POST /api/group
    api.PUT("/group/:group_id", handler.UpdateGroup) // PUT /api/group/:group_id

    api.GET("/directmessage/:user_id", handler.GetDirectMessages) // GET /api/directmessage/:user_id
    api.POST("/directmessage/:user_id", handler.AddDirectMessage) // POST /api/directmessage/:user_id

    api.GET("/groupmessage/:group_id", handler.GetGroupMessages) // GET /api/groupmessage/:group_id
    api.POST("/groupmessage/:group_id", handler.AddGroupMessage) // POST /api/groupmessage/:group_id

    api.GET("/healthrecord", handler.GetHealthRecords) // GET /api/healthrecord
    api.POST("/healthrecord", handler.AddHealthRecord) // POST /api/healthrecord
    api.PUT("/healthrecord/:record_id", handler.UpdateHealthRecord) // PUT /api/healthrecord/:record_id

    api.GET("/invitation/friend", handler.GetFriendInvitations) // GET /api/invitation/friend
    api.POST("/invitation/friend/:user_id", handler.AddFriendInvitation) // POST /api/invitation/friend/:user_id
    api.PUT("/invitation/friend/:invitation_id", handler.UpdateFriendInvitationStatus) // PUT /api/invitation/friend/:invitation_id
    
    api.GET("/invitation/group", handler.GetGroupInvitations) // GET /api/invitation/group
    api.POST("/invitation/group", handler.AddGroupInvitation) // POST /api/invitation/group
    api.PUT("/invitation/group/:invitation_id", handler.UpdateGroupInvitationStatus) // PUT /api/invitation/group/:invitation_id
    

    return e
}