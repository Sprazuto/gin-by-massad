package controllers

import (
	"strconv"
	"strings"

	"lke-app/forms"
	"lke-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeKomponenController ...
type LkeKomponenController struct{}

var lkeKomponenModel = new(models.LkeKomponenModel)
var lkeKomponenForm = new(forms.LkeKomponenForm)

// Create creates a new LKE Komponen record
// @Summary Create LKE Komponen
// @Description Create a new LKE (Lembar Kerja Efektivitas) Komponen record
// @Tags LKE Komponen
// @Accept json
// @Produce json
// @Security Bearer
// @Param createForm body forms.CreateLkeKomponenForm true "LKE Komponen data"
// @Success 200 {object} map[string]interface{} "LKE Komponen created"
// @Failure 406 {object} map[string]interface{} "Validation failed"
// @Router /v1/lke-komponen [post]
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

// All retrieves all LKE Komponen records
// @Summary Get all LKE Komponen records
// @Description Get all LKE Komponen records
// @Tags Used in E-Office, LKE Komponen
// @Produce json,xml
// @Security Bearer
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Komponen records list"
// @Failure 406 {object} map[string]interface{} "Could not get LKE Komponen records"
// @Router /v1/lke-komponens [get]
// @Router /v1/lke-komponens/{format} [get]
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

// One retrieves a specific LKE Komponen record by ID
// @Summary Get LKE Komponen by ID
// @Description Get a specific LKE Komponen record by ID
// @Tags LKE Komponen
// @Produce json,xml
// @Security Bearer
// @Param id path int true "LKE Komponen ID"
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Komponen data"
// @Failure 404 {object} map[string]interface{} "LKE Komponen not found or invalid parameter"
// @Router /v1/lke-komponen/{id} [get]
// @Router /v1/lke-komponen/{id}/{format} [get]
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

// Update modifies an existing LKE Komponen record
// @Summary Update LKE Komponen
// @Description Update an existing LKE Komponen record by ID
// @Tags LKE Komponen
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Komponen ID"
// @Param createForm body forms.CreateLkeKomponenForm true "Updated LKE Komponen data"
// @Success 200 {object} map[string]interface{} "LKE Komponen updated"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Validation failed or could not update"
// @Router /v1/lke-komponen/{id} [put]
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

// Delete removes an LKE Komponen record by ID
// @Summary Delete LKE Komponen
// @Description Delete an LKE Komponen record by ID
// @Tags LKE Komponen
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Komponen ID"
// @Success 200 {object} map[string]interface{} "LKE Komponen deleted"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Could not delete"
// @Router /v1/lke-komponen/{id} [delete]
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
