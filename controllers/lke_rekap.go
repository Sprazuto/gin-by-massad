package controllers

import (
	"strconv"
	"strings"

	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeRekapController ...
type LkeRekapController struct{}

var lkeRekapModel = new(models.LkeRekapModel)
var lkeRekapForm = new(forms.LkeRekapForm)

// Create ...
func (ctrl LkeRekapController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateLkeRekapForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeRekapForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := lkeRekapModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "LkeRekap could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeRekap created", "id": id})
}

// All ...
func (ctrl LkeRekapController) All(c *gin.Context) {
	userID := getUserID(c)

	format := c.Param("format")

	results, err := lkeRekapModel.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get lke_rekap records"})
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
func (ctrl LkeRekapController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")
	format := c.Param("format")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := lkeRekapModel.One(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "LkeRekap not found"})
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
func (ctrl LkeRekapController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.UpdateLkeRekapForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeRekapForm.Update(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = lkeRekapModel.Update(userID, getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "LkeRekap could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeRekap updated"})
}

// Delete ...
func (ctrl LkeRekapController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = lkeRekapModel.Delete(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "LkeRekap could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeRekap deleted"})
}

// GetByOPDAndTahun gets a lke_rekap record by id_opd and tahun with its evaluasi children
func (ctrl LkeRekapController) GetByOPDAndTahun(c *gin.Context) {
	userID := getUserID(c)

	idOPD := c.Param("id_opd")
	tahun := c.Param("tahun")

	getIDOPD, err := strconv.ParseInt(idOPD, 10, 64)
	if getIDOPD == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid id_opd parameter"})
		return
	}

	getTahun, err := strconv.Atoi(tahun)
	if err != nil || getTahun <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid tahun parameter - must be a positive integer"})
		return
	}

	data, err := lkeRekapModel.OneWithEvaluasi(userID, getIDOPD, getTahun)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "LkeRekap not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
