package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Mountain struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
}

var (
	mountains = map[int]Mountain{}
	seq       = 0
)

func main() {
	e := echo.New()

	initRouting(e)

	e.Logger.Fatal(e.Start(":1313"))
}

func initRouting(e *echo.Echo) {
	e.GET("/api/v1/mountains/:id", HandleAPIGetMountain)
	e.POST("/api/v1/mountains", HandleAPISetMountain)
}

func HandleAPIGetMountain(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("errors when id convert to int: %s", c.Param("id"))
	}
	return c.JSON(http.StatusOK, mountains[id])
}

func HandleAPISetMountain(c echo.Context) error {
	param := new(Mountain)
	if err := c.Bind(param); err != nil {
		return err
	}
	mountain := Mountain{
		ID:     seq,
		Name:   param.Name,
		Height: param.Height,
	}

	mountains[seq] = mountain
	seq++

	return c.JSON(http.StatusOK, mountain)
}
