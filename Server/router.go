package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/signup", handler.Signup) // POST /signup
    e.POST("/login", handler.Login) // POST /login

    return e
}