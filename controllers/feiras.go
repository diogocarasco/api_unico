package controllers

import (
	"api_unico/database"
	"api_unico/models"
	"api_unico/models/input"
	"api_unico/utils"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetFeiras retrieve all registers from FEIRAS
func GetFeiras(c *gin.Context) {

	var feiras []models.Feiras

	type query struct {
		distrito   string
		regiao5    string
		nome_feira string
		bairro     string
	}

	var querystring = models.Feiras{}
	bairro := c.Query("bairro")
	if bairro != "" {
		querystring.BAIRRO = bairro
		database.DB.Where(&querystring, "BAIRRO").Find(&feiras)
	} else {
		database.DB.Find(&feiras)
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(feiras)
}

// GetFeirasById retrieve registers from FEIRAS by ID
func GetFeirasById(c *gin.Context) {

	var feiras *models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		logrus.Info("data:" + "[]")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(feiras)

}

// InsertFeiras create a new register into FEIRAS
func InsertFeiras(c *gin.Context) {

	var feirasinput input.CreateFeirasInput
	var feiras models.Feiras

	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Info("error:" + err.Error())
		return
	}

	database.DB.Create(&feiras)
	c.JSON(http.StatusOK, gin.H{"id": feiras.ID})
	logrus.Debug(feiras.ID)

}

// DeleteFeiras remove a FEIRAS register by ID
func DeleteFeiras(c *gin.Context) {

	var feiras models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		return
	}

	database.DB.Delete(&feiras)

	c.JSON(http.StatusOK, gin.H{"data": true})
	logrus.Debug(gin.H{"data": true})

}

// UpdateFeiras update row from FEIRAS
func UpdateFeiras(c *gin.Context) {

	var feirasinput models.Feiras
	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Info("error:" + err.Error())
		return
	}

	var feiras models.Feiras
	err = database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		logrus.Info("error:" + err.Error())
		return
	}

	upderr := database.DB.Model(&feiras).Omit("id").Updates(feirasinput)
	if upderr != nil {
		logrus.Debug(upderr.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(gin.H{"data": feiras})

}

func ImportFeiras(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		logrus.Debug(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "data/import/"+filename); err != nil {
		logrus.Debug(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
		return
	}

	content, errors := utils.ImportFile("data/import/" + filename)
	if errors != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
	} else {
		err := database.DB.Create(&content)
		if err.Error != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
			logrus.Debug(err.Error)
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("File %s processed successfully.", file.Filename))
	}

}
