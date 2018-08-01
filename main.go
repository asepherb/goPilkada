package main

import (
	"github.com/labstack/echo"
)


func main() {
	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "REST API PILKADA") })

	e.Logger.Fatal(e.Start(":1323"))
}
