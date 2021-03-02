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
    api.GET("/directmessage/:with", handler.GetDirectMessages) // GET /api/directmessage/:with
    api.POST("/directmessage/:with", handler.AddDirectMessage) // POST /api/directmessage/:with
    api.GET("/groupmessage/:group_id", handler.GetGroupMessages) // GET /api/groupmessage/:group_id
    api.POST("/groupmessage/:group_id", handler.AddGroupMessage) // POST /api/groupmessage/:group_id

    return e
}