package infrastructure

import (
	"github.com/labstack/echo/v4"

	"github.com/shiki-tak/YamaWikipedia/interfaces/controllers"
)

var Router *echo.Echo

func init() {
	e := echo.New()
	mountainController := controllers.NewMountainController(NewLevelDBHandler())

	api := e.Group("/api")
	api.GET("/v1/mountains/:id", func(c echo.Context) error { return mountainController.Show(c) })
	api.POST("/v1/mountains", func(c echo.Context) error { return mountainController.Create(c) })

	Router = e
}
