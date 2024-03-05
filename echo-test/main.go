package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	sort := c.QueryParam("sort")
	limit := c.QueryParam("limit")

	res := map[string]interface{}{
		"sort":  sort,
		"limit": limit,
	}
	return c.JSON(http.StatusOK, res)
}
