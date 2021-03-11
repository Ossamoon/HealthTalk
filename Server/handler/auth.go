package handler


import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
	"github.com/Ossamoon/HealthTalk/Server/model"
)


type (
    jwtCustomClaims struct {
        UID  uint   `json:"uid"`
        Name string `json:"name"`
        jwt.StandardClaims
    }

    SignUpResponse struct {
        ID          uint
        CreatedAt   time.Time
        Name        string      `json:"name"`
        Email       string      `json:"email"`
    }

    LoginResponse struct {
        Token       string      `json:"token"`
    }
)

var (
    signingKey = []byte("secretKey")

    Config = middleware.JWTConfig{
        Claims:     &jwtCustomClaims{},
        SigningKey: signingKey,
    }
)


func Signup(c echo.Context) error {
    user := new(model.User)
    if err := c.Bind(user); err != nil {
        return err
    }

    if user.Name == "" || user.Password == "" || user.Email == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid name or password or email",
        }
    }

    if u := model.FindUser(&model.User{Email: user.Email}); u.ID != 0 {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "email already exists",
        }
    }

    model.CreateUser(user)
    responce := SignUpResponse {
        ID: user.Model.ID,
        CreatedAt: user.Model.CreatedAt,
        Name: user.Name,
        Email: user.Email,
    }

    return c.JSON(http.StatusCreated, responce)
}


func Login(c echo.Context) error {
    u := new(model.User)
    if err := c.Bind(u); err != nil {
        return err
    }

    user := model.FindUser(&model.User{Name: u.Name})
    if user.ID == 0 || user.Password != u.Password {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid name or password",
        }
    }

    claims := &jwtCustomClaims{
        user.Model.ID,
        user.Name,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString(signingKey)
    if err != nil {
        return err
    }

    responce := LoginResponse {
        Token: t,
    }

    return c.JSON(http.StatusOK, responce)
}


func userIDFromToken(c echo.Context) uint {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtCustomClaims)
    uid := claims.UID
    return uid
}