package controllers

import (
	"strconv"
	"strings"
	"time"

	"lke-app/forms"
	"lke-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// LkeRekapController ...
type LkeRekapController struct{}

var lkeRekapModel = new(models.LkeRekapModel)
var lkeRekapForm = new(forms.LkeRekapForm)

// Create creates a new LKE Rekap record
// @Summary Create LKE Rekap
// @Description Create a new LKE (Lembar Kerja Efektivitas) Rekap record for the authenticated user
// @Tags LKE Rekap
// @Accept json
// @Produce json
// @Security Bearer
// @Param createForm body forms.CreateLkeRekapForm true "LKE Rekap data"
// @Success 200 {object} map[string]interface{} "LKE Rekap created"
// @Failure 406 {object} map[string]interface{} "Validation failed"
// @Router /v1/lke-rekap [post]
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

// All retrieves all LKE Rekap records for the authenticated user
// @Summary Get all LKE Rekap records
// @Description Get all LKE Rekap records for the authenticated user, optionally filtered by year
// @Tags Used in E-Office, LKE Rekap
// @Produce json,xml
// @Security Bearer
// @Param tahun path int false "Year filter (defaults to current year - 1)"
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Rekap records list"
// @Failure 400 {object} map[string]interface{} "Invalid year parameter"
// @Failure 406 {object} map[string]interface{} "Could not get LKE Rekap records"
// @Router /v1/lke-rekaps [get]
// @Router /v1/lke-rekaps/{tahun} [get]
// @Router /v1/lke-rekaps/{tahun}/{format} [get]
func (ctrl LkeRekapController) All(c *gin.Context) {
	userID := getUserID(c)

	format := c.Param("format")
	year := c.Param("year")

	var tahun int
	if year == "" {
		tahun = time.Now().Year() - 1 // Default to current tahun - 1
	} else {
		var err error
		tahun, err = strconv.Atoi(year)
		if err != nil || tahun <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid year parameter - must be a positive integer"})
			return
		}
	}

	results, err := lkeRekapModel.All(userID, tahun)
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

// One retrieves a specific LKE Rekap record by ID
// @Summary Get LKE Rekap by ID
// @Description Get a specific LKE Rekap record by ID for the authenticated user
// @Tags Used in E-Office, LKE Rekap
// @Produce json,xml
// @Security Bearer
// @Param id path int true "LKE Rekap ID"
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Rekap data"
// @Failure 404 {object} map[string]interface{} "LKE Rekap not found or invalid parameter"
// @Router /v1/lke-rekap/{id} [get]
// @Router /v1/lke-rekap/{id}/{format} [get]
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

// Update modifies an existing LKE Rekap record
// @Summary Update LKE Rekap
// @Description Update an existing LKE Rekap record by ID
// @Tags Used in E-Office, LKE Rekap
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Rekap ID"
// @Param updateForm body forms.UpdateLkeRekapForm true "Updated LKE Rekap data"
// @Success 200 {object} map[string]interface{} "LKE Rekap updated"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Validation failed or could not update"
// @Router /v1/lke-rekap/{id} [put]
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

// Delete removes an LKE Rekap record by ID
// @Summary Delete LKE Rekap
// @Description Delete an LKE Rekap record by ID
// @Tags LKE Rekap
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Rekap ID"
// @Success 200 {object} map[string]interface{} "LKE Rekap deleted"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Could not delete"
// @Router /v1/lke-rekap/{id} [delete]
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

// GetByOPDAndTahun gets a LKE Rekap record by OPD ID and year with its evaluasi children
// @Summary Get LKE Rekap by OPD and Year
// @Description Get a LKE Rekap record by OPD ID and year, including its evaluasi children records
// @Tags Used in E-Office, LKE Rekap
// @Produce json
// @Security Bearer
// @Param id_opd path int true "OPD ID"
// @Param tahun path int true "Year"
// @Success 200 {object} map[string]interface{} "LKE Rekap data with evaluasi children"
// @Failure 404 {object} map[string]interface{} "LKE Rekap not found or invalid parameters"
// @Router /v1/lke-rekap/opd/{id_opd}/tahun/{tahun} [get]
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
