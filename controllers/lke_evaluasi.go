package controllers

import (
	"strconv"
	"strings"

	"lke-app/forms"
	"lke-app/models"
	"lke-app/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LkeEvaluasiController struct{}

var lkeEvaluasiModel = new(models.LkeEvaluasiModel)
var lkeEvaluasiForm = new(forms.LkeEvaluasiForm)

// SyncEvaluasi duplicates jawaban into evaluasi for given lke_rekap_id
// @Summary Sync LKE Evaluasi
// @Description Synchronize jawaban data into evaluasi records for the specified LKE Rekap ID
// @Tags Used in E-Office, LKE Evaluasi
// @Produce json
// @Security Bearer
// @Param lke_rekap_id query int true "LKE Rekap ID"
// @Success 200 {object} map[string]interface{} "Successfully synced evaluasi from jawaban"
// @Failure 400 {object} map[string]interface{} "Invalid lke_rekap_id parameter"
// @Failure 500 {object} map[string]interface{} "Failed to sync evaluasi or update rekap values"
// @Router /v1/lke-evaluasi/sync-evaluasi [get]
func (ctrl LkeEvaluasiController) SyncEvaluasi(c *gin.Context) {
	lkeRekapIDStr := c.Query("lke_rekap_id")
	if lkeRekapIDStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "lke_rekap_id parameter is required"})
		return
	}

	lkeRekapID, err := strconv.ParseInt(lkeRekapIDStr, 10, 64)
	if err != nil || lkeRekapID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid lke_rekap_id parameter"})
		return
	}

	err = lkeEvaluasiModel.SyncJawabanToEvaluasi(lkeRekapID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to sync evaluasi", "error": err.Error()})
		return
	}

	err = lkeEvaluasiModel.UpdateRekapValues(lkeRekapID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update rekap values", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully synced evaluasi from jawaban"})
}

// Create creates or updates an LKE Evaluasi record
// @Summary Create or Update LKE Evaluasi
// @Description Create a new LKE Evaluasi record or update existing one. Supports file upload for PUT requests.
// @Tags Used in E-Office, LKE Evaluasi
// @Accept json,multipart/form-data
// @Produce json
// @Security Bearer
// @Param createForm body forms.CreateLkeEvaluasiForm true "LKE Evaluasi data"
// @Param file formData file false "File to upload (for PUT requests only)"
// @Success 200 {object} map[string]interface{} "LKE Evaluasi created or updated"
// @Failure 400 {object} map[string]interface{} "File required for PUT or invalid request"
// @Failure 406 {object} map[string]interface{} "Validation failed"
// @Failure 500 {object} map[string]interface{} "Failed to upload file or update rekap values"
// @Router /v1/lke-evaluasi [post]
// @Router /v1/lke-evaluasi [put]
func (ctrl LkeEvaluasiController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateLkeEvaluasiForm

	if validationErr := c.ShouldBind(&form); validationErr != nil {
		message := lkeEvaluasiForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	// Handle file upload only for PUT method
	if c.Request.Method == http.MethodPut {
		file, err := c.FormFile("file") // Ensure the file is sent with the key "file"
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "File is required"})
			return
		}

		// Upload file to MinIO
		fileURL, err := services.UploadFile(file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload file", "error": err.Error()})
			return
		}
		form.Berkas = &fileURL
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

// All retrieves all LKE Evaluasi records for the authenticated user
// @Summary Get all LKE Evaluasi records
// @Description Get all LKE Evaluasi records for the authenticated user
// @Tags LKE Evaluasi
// @Produce json,xml
// @Security Bearer
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Evaluasi records list"
// @Failure 406 {object} map[string]interface{} "Could not get LKE Evaluasi records"
// @Router /v1/lke-evaluasis [get]
// @Router /v1/lke-evaluasis/{format} [get]
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

// One retrieves a specific LKE Evaluasi record by ID
// @Summary Get LKE Evaluasi by ID
// @Description Get a specific LKE Evaluasi record by ID for the authenticated user
// @Tags LKE Evaluasi
// @Produce json,xml
// @Security Bearer
// @Param id path int true "LKE Evaluasi ID"
// @Param format path string false "Response format: json or xml (defaults to json)"
// @Success 200 {object} map[string]interface{} "LKE Evaluasi data"
// @Failure 404 {object} map[string]interface{} "LKE Evaluasi not found or invalid parameter"
// @Router /v1/lke-evaluasi/{id} [get]
// @Router /v1/lke-evaluasi/{id}/{format} [get]
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

// Update modifies an existing LKE Evaluasi record
// @Summary Update LKE Evaluasi
// @Description Update an existing LKE Evaluasi record by ID
// @Tags LKE Evaluasi
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Evaluasi ID"
// @Param createForm body forms.CreateLkeEvaluasiForm true "Updated LKE Evaluasi data"
// @Success 200 {object} map[string]interface{} "LKE Evaluasi updated"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Validation failed or could not update"
// @Router /v1/lke-evaluasi/{id} [put]
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

// Delete removes an LKE Evaluasi record by ID
// @Summary Delete LKE Evaluasi
// @Description Delete an LKE Evaluasi record by ID
// @Tags LKE Evaluasi
// @Produce json
// @Security Bearer
// @Param id path int true "LKE Evaluasi ID"
// @Success 200 {object} map[string]interface{} "LKE Evaluasi deleted"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Could not delete"
// @Router /v1/lke-evaluasi/{id} [delete]
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

// GetSignedURL returns a signed URL for the specified LkeRekapID and KodeEvaluasi
// @Summary Get signed URL for LKE Evaluasi file
// @Description Generate a signed URL to access a file associated with an LKE Evaluasi record
// @Tags Used in E-Office, LKE Evaluasi
// @Produce json
// @Security Bearer
// @Param lke_rekap_id path int true "LKE Rekap ID"
// @Param kode_evaluasi path string true "Kode Evaluasi"
// @Success 200 {object} map[string]interface{} "Signed URL for file access"
// @Failure 400 {object} map[string]interface{} "Missing required parameters"
// @Failure 404 {object} map[string]interface{} "Record not found"
// @Failure 500 {object} map[string]interface{} "Failed to generate signed URL"
// @Router /v1/lke-evaluasi/signed-url/{lke_rekap_id}/{kode_evaluasi} [get]
func (ctrl LkeEvaluasiController) GetSignedURL(c *gin.Context) {
	lkeRekapID := c.Param("lke_rekap_id")
	kodeEvaluasi := c.Param("kode_evaluasi")

	if lkeRekapID == "" || kodeEvaluasi == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "lke_rekap_id and kode_evaluasi are required"})
		return
	}

	// Logic to retrieve the record and generate the signed URL
	lkeRekapIDInt, err := strconv.ParseInt(lkeRekapID, 10, 64)
	record, err := lkeEvaluasiModel.OneByRekapIDAndKodeEvaluasi(getUserID(c), lkeRekapIDInt, kodeEvaluasi)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	signedURL, err := services.GenerateSignedURL(record.Berkas.String)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate signed URL", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"signed_url": signedURL})
}
