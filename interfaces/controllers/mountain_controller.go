package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/shiki-tak/YamaWikipedia/domain"
	"github.com/shiki-tak/YamaWikipedia/interfaces/database"
	"github.com/shiki-tak/YamaWikipedia/usecase"
)

type MountainController struct {
	Interactor usecase.MountainInteractor
}

func NewMountainController(levelDBhandler database.LevelDBHandler) *MountainController {
	return &MountainController{
		Interactor: usecase.MountainInteractor{
			MountainRepository: &database.MountainRepository{
				LevelDBHandler: levelDBhandler,
			},
		},
	}
}

func (controller *MountainController) Create(c echo.Context) error {
	m := domain.Mountain{}
	err := c.Bind(&m)
	if err != nil {
		return c.JSON(500, fmt.Errorf("api request error:%s", err))
	}
	err = controller.Interactor.Add(m)
	if err != nil {
		return c.JSON(500, fmt.Errorf("api request error:%s", err))
	}
	return c.JSON(201, nil)
}

func (controller *MountainController) Show(c echo.Context) error {
	id := c.Param("id")
	mountain, err := controller.Interactor.MountainById(id)
	if err != nil {
		return c.JSON(500, fmt.Errorf("api request error:%s", err))
	}
	return c.JSON(200, mountain)
}

func (controller *MountainController) Index(c echo.Context) error {
	mountains, err := controller.Interactor.AllMountains()
	if err != nil {
		return c.JSON(500, fmt.Errorf("api request error:%s", err))
	}
	return c.JSON(200, mountains)
}
