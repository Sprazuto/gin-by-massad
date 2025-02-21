package controllers

import (
	"strconv"
	"strings"

	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeEvaluasiController ...
type LkeEvaluasiController struct{}

var lkeEvaluasiModel = new(models.LkeEvaluasiModel)
var lkeEvaluasiForm = new(forms.LkeEvaluasiForm)

// Create ...
func (ctrl LkeEvaluasiController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateLkeEvaluasiForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeEvaluasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := lkeEvaluasiModel.CreateOrUpdate(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "LkeEvaluasi could not be created or updated"})
		return
	}

	// Update kelengkapan and nilai_capaian in lke_rekap
	err = lkeEvaluasiModel.UpdateRekapValues(form.LkeRekapID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update lke_rekap values",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeEvaluasi created or updated", "id": id})

}

// All ...
func (ctrl LkeEvaluasiController) All(c *gin.Context) {
	userID := getUserID(c)

	format := c.Param("format")

	results, err := lkeEvaluasiModel.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get lke_evaluasi records"})
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
func (ctrl LkeEvaluasiController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")
	format := c.Param("format")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := lkeEvaluasiModel.One(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "LkeEvaluasi not found"})
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
func (ctrl LkeEvaluasiController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.CreateLkeEvaluasiForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeEvaluasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = lkeEvaluasiModel.Update(userID, getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "LkeEvaluasi could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeEvaluasi updated"})
}

// Delete ...
func (ctrl LkeEvaluasiController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = lkeEvaluasiModel.Delete(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "LkeEvaluasi could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LkeEvaluasi deleted"})
}
