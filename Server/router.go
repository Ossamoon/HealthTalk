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
    api.GET("/directmessage", handler.GetDirectMessages) // GET /api/directmessage
    api.POST("/directmessage", handler.AddDirectMessage) // POST /api/directmessage

    return e
}