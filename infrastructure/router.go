package infrastructure

import (
	"github.com/labstack/echo/v4"

	"github.com/shiki-tak/YamaWikipedia/interfaces/controllers"
)

func Init() {
	e := echo.New()
	mountainController := controllers.NewMountainController(NewLevelDBHandler())

	api := e.Group("/api")
	api.GET("/v1/mountains/:id", func(c echo.Context) error { return mountainController.Show(c) })
	api.GET("/v1/mountains", func(c echo.Context) error { return mountainController.Index(c) })
	api.POST("/v1/mountains", func(c echo.Context) error { return mountainController.Create(c) })

	e.Logger.Fatal(e.Start(":1313"))
}
