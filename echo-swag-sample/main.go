package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	Response struct {
		Int64  int64  `json:"int64"`
		String string `json:"string"`
		World  *Item  `json:"world"`
	}

	Item struct {
		Text string `json:"text"`
	}
)

// @title         example
// @version       1.0
// @license.name  license name
// @BasePath      /
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}

// hello godoc
// @Summary  Hello World !
// @ID       HelloWorldIndex
// @Tags     HelloWorld
// @Produce  json
// @Success  200  {object}  Response
// @Router   / [get]
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{
		Int64:  1,
		String: "example",
		World: &Item{
			Text: "hello world !",
		},
	})
}
