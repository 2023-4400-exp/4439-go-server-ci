package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/", Hello())
    e.GET("/api/", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "4439: GoはGoogleが開発した言語 ver.2")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4439", "message": "GoはGoogleが開発した言語"})
    }
}
