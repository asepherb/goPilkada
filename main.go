package main

import (
	"net/http"
	"github.com/asepherb/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "REST API Pemilu!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}