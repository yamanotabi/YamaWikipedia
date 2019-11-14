package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Mountain struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
}

var (
	mountains = map[int]Mountain{}
	seq       = 0
	db        *leveldb.DB
)

func main() {
	e := echo.New()

	// open db
	db, _ = leveldb.OpenFile("./db", nil)
	defer db.Close()

	initRouting(e)

	e.Logger.Fatal(e.Start(":1313"))
}

func initRouting(e *echo.Echo) {
	e.GET("/api/v1/mountains/:id", HandleAPIGetMountain)
	e.POST("/api/v1/mountains", HandleAPISetMountain)
}

func HandleAPIGetMountain(c echo.Context) error {
	jsonBytes, err := db.Get([]byte(c.Param("id")), nil)
	if err != nil {
		return fmt.Errorf("Get Data error:", err)
	}
	mountain := new(Mountain)
	err = json.Unmarshal(jsonBytes, mountain)
	if err != nil {
		return fmt.Errorf("JSON Unmarshal error:", err)
	}
	return c.JSON(http.StatusOK, mountain)
}

func HandleAPISetMountain(c echo.Context) error {
	param := new(Mountain)
	if err := c.Bind(param); err != nil {
		return fmt.Errorf("param bind error:", err)
	}
	mountain := Mountain{
		ID:     seq,
		Name:   param.Name,
		Height: param.Height,
	}
	seqStr := strconv.Itoa(seq)

	jsonBytes, err := json.Marshal(mountain)
	if err != nil {
		return fmt.Errorf("JSON Marshal error:", err)
	}

	err = db.Put([]byte(seqStr), jsonBytes, nil)
	if err != nil {
		return fmt.Errorf("Put Data error:", err)
	}
	seq++

	return c.JSON(http.StatusOK, mountain)
}
