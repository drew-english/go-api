package handlers

import (
	"go-api/internal/models/app_setting"
	"github.com/gin-gonic/gin"
)

type AppSettingHandler struct {
	AppSettings *app_setting.AppSettings
}

func (as *AppSettingHandler) Show(c *gin.Context) {
	name := c.Param("name")

	// get value from db
	c.JSON(200, map[string]string{ "name": name })
}

func (as *AppSettingHandler) Create(c *gin.Context) {

}

func (as *AppSettingHandler) Update(c *gin.Context) {
	name := c.Param("name")

	// update value in db
	c.JSON(200, map[string]string{ "name": name })
}

func (as *AppSettingHandler) Delete(c *gin.Context) {
	name := c.Param("name")

	// delete value from db
	c.JSON(200, map[string]string{ "name": name })
}
