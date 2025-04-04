package controllers

import (
	"strconv"
	"strings"

	"lke-app/forms"
	"lke-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeRekomendasiController ...
type LkeRekomendasiController struct{}

var lkeRekomendasiModel = new(models.LkeRekomendasiModel)
var lkeRekomendasiForm = new(forms.LkeRekomendasiForm)

// Create ...
func (ctrl LkeRekomendasiController) Create(c *gin.Context) {
	var form forms.CreateLkeRekomendasiForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeRekomendasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := lkeRekomendasiModel.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Rekomendasi could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rekomendasi created", "id": id})
}

// All ...
func (ctrl LkeRekomendasiController) All(c *gin.Context) {
	format := c.Param("format")

	results, err := lkeRekomendasiModel.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get rekomendasi"})
		return
	}

	switch strings.ToLower(format) {
	case "json":
		c.JSON(http.StatusOK, gin.H{"results": results})
	case "xml":
		c.XML(http.StatusOK, gin.H{"results": results})
	case "yaml":
		c.YAML(http.StatusOK, gin.H{"results": results})
	default:
		c.JSON(http.StatusOK, gin.H{"results": results})
	}
}

// One ...
func (ctrl LkeRekomendasiController) One(c *gin.Context) {
	id := c.Param("id")
	format := c.Param("format")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := lkeRekomendasiModel.One(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Rekomendasi not found"})
		return
	}

	switch strings.ToLower(format) {
	case "json":
		c.JSON(http.StatusOK, gin.H{"data": data})
	case "xml":
		c.XML(http.StatusOK, gin.H{"data": data})
	case "yaml":
		c.YAML(http.StatusOK, gin.H{"data": data})
	default:
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// Update ...
func (ctrl LkeRekomendasiController) Update(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.CreateLkeRekomendasiForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeRekomendasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = lkeRekomendasiModel.Update(getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Rekomendasi could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rekomendasi updated"})
}

// Delete ...
func (ctrl LkeRekomendasiController) Delete(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = lkeRekomendasiModel.Delete(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Rekomendasi could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rekomendasi deleted"})
}
