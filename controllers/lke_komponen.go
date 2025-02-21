package controllers

import (
	"strconv"
	"strings"

	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeKomponenController ...
type LkeKomponenController struct{}

var lkeKomponenModel = new(models.LkeKomponenModel)
var lkeKomponenForm = new(forms.LkeKomponenForm)

// Create ...
func (ctrl LkeKomponenController) Create(c *gin.Context) {
	var form forms.CreateLkeKomponenForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeKomponenForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := lkeKomponenModel.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Komponen could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Komponen created", "id": id})
}

// All ...
func (ctrl LkeKomponenController) All(c *gin.Context) {
	format := c.Param("format")

	results, err := lkeKomponenModel.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get komponen"})
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
func (ctrl LkeKomponenController) One(c *gin.Context) {
	id := c.Param("id")
	format := c.Param("format")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := lkeKomponenModel.One(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Komponen not found"})
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
func (ctrl LkeKomponenController) Update(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.CreateLkeKomponenForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeKomponenForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = lkeKomponenModel.Update(getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Komponen could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Komponen updated"})
}

// Delete ...
func (ctrl LkeKomponenController) Delete(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = lkeKomponenModel.Delete(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Komponen could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Komponen deleted"})
}
