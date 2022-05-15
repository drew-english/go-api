package handlers

import (
	"errors"

	"go-api/internal/models/app_setting"

	"github.com/gin-gonic/gin"
)

type AppSettingHandler struct {
	AppSettings *app_setting.AppSettings
}

// @Description Returns the value for given name
// @Tags App Settings
// @Produce json
// @Param name path string true "Name of the App Setting"
// @Success 200 {object} app_setting.AppSetting
// @Failure 404 {object} docutil.NotFoundResponse
// @Failure 500 {object} nil
// @Router /v1/app_setting/{name} [get]
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

// @Description Creates an App Setting with given name and value
// @Tags App Settings
// @Accept json
// @Produce json
// @Param body body app_setting.AppSetting true "App Setting data"
// @Success 201 {object} nil
// @Failure 500 {object} nil
// @Router /v1/app_setting [post]
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

// @Description Updates the App Setting matching a given namm
// @Tags App Settings
// @Accept json
// @Produce json
// @Param body body app_setting.AppSetting true "App Setting with new value"
// @Success 200 {object} nil "Sucessfully updated"
// @Failure 404 {object} docutil.NotFoundResponse
// @Failure 500 {object} nil
// @Router /v1/app_setting [put]
func (ash *AppSettingHandler) Update(c *gin.Context) {

	// update value in db
	c.Status(200)
}

// @Description Deletes the App Setting matching a given namm
// @Tags App Settings
// @Accept json
// @Produce json
// @Param name path string true "App Setting to delete"
// @Success 200 {object} nil "Sucessfully deleted"
// @Failure 404 {object} docutil.NotFoundResponse
// @Failure 500 {object} nil
// @Router /v1/app_setting/{name} [delete]
func (ash *AppSettingHandler) Delete(c *gin.Context) {
	name := c.Param("name")

	// delete value from db
	c.JSON(200, map[string]string{ "name": name })
}
