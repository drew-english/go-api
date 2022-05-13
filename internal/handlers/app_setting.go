package handlers

import (
	"errors"

	"go-api/internal/models/app_setting"

	"github.com/gin-gonic/gin"
)

type AppSettingHandler struct {
	AppSettings *app_setting.AppSettings
}

func (ash *AppSettingHandler) Show(c *gin.Context) {
	name := c.Param("name")

	foundSetting, err := ash.AppSettings.Find(name)
	if err != nil {
		if errors.Is(err, app_setting.NOT_FOUND_ERROR{}) {
			c.JSON(404, map[string]string { "message": "Not found" })
			return
		}

		panic("Unknown error fetching app setting: " + name)
	}
	c.JSON(200, foundSetting)
}

func (ash *AppSettingHandler) Create(c *gin.Context) {
	var newSetting app_setting.AppSetting

	err := c.BindJSON(&newSetting)
	if err != nil {
		panic("Error binding app setting")
	}

	err = ash.AppSettings.Create(&newSetting)
	if err != nil {
		panic("Error creating app setting")
	}

	c.Status(201)
}

func (ash *AppSettingHandler) Update(c *gin.Context) {
	name := c.Param("name")

	// update value in db
	c.JSON(200, map[string]string{ "name": name })
}

func (ash *AppSettingHandler) Delete(c *gin.Context) {
	name := c.Param("name")

	// delete value from db
	c.JSON(200, map[string]string{ "name": name })
}
