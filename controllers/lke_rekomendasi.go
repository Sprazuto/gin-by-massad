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

// Create creates or updates an LKE Rekomendasi record
// @Summary Create or Update LKE Rekomendasi
// @Description Create a new LKE Rekomendasi record or update existing one based on parent_id
// @Tags Used in E-Office, LKE Rekomendasi
// @Accept json
// @Produce json
// @Security Bearer
// @Param createForm body forms.CreateLkeRekomendasiForm true "LKE Rekomendasi data"
// @Success 200 {object} map[string]interface{} "LKE Rekomendasi created or updated"
// @Failure 406 {object} map[string]interface{} "Validation failed"
// @Router /v1/lke-rekomendasi [post]
func (ctrl LkeRekomendasiController) Create(c *gin.Context) {
	var form forms.CreateLkeRekomendasiForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := lkeRekomendasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, isUpdate, err := lkeRekomendasiModel.Upsert(form)
	if err != nil {
		if isUpdate {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Rekomendasi could not be updated"})
		} else {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Rekomendasi could not be created"})
		}
		return
	}

	if isUpdate {
		c.JSON(http.StatusOK, gin.H{"message": "Rekomendasi updated", "id": id})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Rekomendasi created", "id": id})
	}
}

// All retrieves all LKE Rekomendasi records
// @Summary Get all LKE Rekomendasi records
// @Description Get all LKE Rekomendasi records
// @Tags LKE Rekomendasi
// @Produce json,xml
// @Security Bearer
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Rekomendasi records list"
// @Failure 406 {object} map[string]interface{} "Could not get LKE Rekomendasi records"
// @Router /v1/lke-rekomendasis [get]
// @Router /v1/lke-rekomendasis/{format} [get]
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

// One retrieves a specific LKE Rekomendasi record by ID
// @Summary Get LKE Rekomendasi by ID
// @Description Get a specific LKE Rekomendasi record by ID
// @Tags LKE Rekomendasi
// @Produce json,xml
// @Security Bearer
// @Param id path int true "LKE Rekomendasi ID"
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Rekomendasi data"
// @Failure 404 {object} map[string]interface{} "LKE Rekomendasi not found or invalid parameter"
// @Router /v1/lke-rekomendasi/{id} [get]
// @Router /v1/lke-rekomendasi/{id}/{format} [get]
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

// Update modifies an existing LKE Rekomendasi record
// @Summary Update LKE Rekomendasi
// @Description Update an existing LKE Rekomendasi record by ID
// @Tags LKE Rekomendasi
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Rekomendasi ID"
// @Param createForm body forms.CreateLkeRekomendasiForm true "Updated LKE Rekomendasi data"
// @Success 200 {object} map[string]interface{} "LKE Rekomendasi updated"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Validation failed or could not update"
// @Router /v1/lke-rekomendasi/{id} [put]
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

// Delete removes an LKE Rekomendasi record by ID
// @Summary Delete LKE Rekomendasi
// @Description Delete an LKE Rekomendasi record by ID
// @Tags LKE Rekomendasi
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Rekomendasi ID"
// @Success 200 {object} map[string]interface{} "LKE Rekomendasi deleted"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Could not delete"
// @Router /v1/lke-rekomendasi/{id} [delete]
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
