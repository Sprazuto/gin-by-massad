package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"aset-app/forms"
	"aset-app/models"
	"aset-app/services"

	"github.com/gin-gonic/gin"
)

type VehicleAssetController struct{}

// Helper function to convert sql.NullInt64 to simple JSON value
func nullInt64ToJSON(n sql.NullInt64) interface{} {
	if n.Valid {
		return n.Int64
	}
	return nil
}

// Create creates a new vehicle asset
func (v *VehicleAssetController) Create(c *gin.Context) {
	var form forms.CreateVehicleAssetForm

	// Handle JSON requests
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	asset := models.VehicleAsset{
		LicensePlate:    form.LicensePlate,
		STNKStatus:      form.STNKStatus,
		BPKBNumber:      form.BPKBNumber,
		BPKBStatus:      form.BPKBStatus,
		ChassisNumber:   form.ChassisNumber,
		MachineNumber:   form.MachineNumber,
		TypeName:        form.TypeName,
		WheelsID:        sql.NullInt64{Int64: form.WheelsID.Int64, Valid: form.WheelsID.Valid},
		ModelID:         sql.NullInt64{Int64: form.ModelID.Int64, Valid: form.ModelID.Valid},
		ColorID:         sql.NullInt64{Int64: form.ColorID.Int64, Valid: form.ColorID.Valid},
		FuelID:          sql.NullInt64{Int64: form.FuelID.Int64, Valid: form.FuelID.Valid},
		OwningID:        sql.NullInt64{Int64: form.OwningID.Int64, Valid: form.OwningID.Valid},
		BrandID:         sql.NullInt64{Int64: form.BrandID.Int64, Valid: form.BrandID.Valid},
		CCCapacity:      sql.NullInt64{Int64: form.CCCapacity.Int64, Valid: form.CCCapacity.Valid},
		ManufactureYear: int(form.ManufactureYear),
		TaxDueDate:      sql.NullInt64{Int64: form.TaxDueDate.Int64, Valid: form.TaxDueDate.Valid},
		CurrentOwner:    form.CurrentOwner,
		CompanyID:       int(form.CompanyID.Int64),
		Status:          form.Status,
		CreatedBy:       int(userID),
	}

	id, err := asset.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetAll returns all vehicle assets
func (v *VehicleAssetController) GetAll(c *gin.Context) {
	assets, err := models.GetAllVehicleAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format response for each asset
	var response []gin.H
	for _, asset := range assets {
		response = append(response, gin.H{
			"id":                           asset.ID,
			"license_plate":                asset.LicensePlate,
			"stnk_status":                  asset.STNKStatus,
			"bpkb_number":                  asset.BPKBNumber,
			"bpkb_status":                  asset.BPKBStatus,
			"chassis_number":               asset.ChassisNumber,
			"machine_number":               asset.MachineNumber,
			"cc_capacity":                  nullInt64ToJSON(asset.CCCapacity),
			"manufacture_year":             asset.ManufactureYear,
			"tax_due_date":                 nullInt64ToJSON(asset.TaxDueDate),
			"last_tax_payment_date":        nullInt64ToJSON(asset.LastTaxPaymentDate),
			"current_owner":                asset.CurrentOwner,
			"stnk_photo_path":              asset.STNKPhotoPath,
			"stnk_photo_verified":          asset.STNKPhotoVerified,
			"bpkb_photo_path":              asset.BPKBPhotoPath,
			"bpkb_photo_verified":          asset.BPKBPhotoVerified,
			"owner_id_photo_path":          asset.OwnerIDPhotoPath,
			"owner_id_photo_verified":      asset.OwnerIDPhotoVerified,
			"vehicle_photo_path":           asset.VehiclePhotoPath,
			"vehicle_photo_verified":       asset.VehiclePhotoVerified,
			"payment_billing_photo_path":   asset.PaymentBillingPhotoPath,
			"payment_billing_date":         nullInt64ToJSON(asset.PaymentBillingDate),
			"recommendation_document_path": asset.RecommendationDocPath,
			"e_sign_status":                asset.ESignStatus,
			"notes":                        asset.Notes,
			"status":                       asset.Status,
			"created_at":                   asset.CreatedAt,
			"updated_at":                   asset.UpdatedAt,
			"created_by":                   asset.CreatedBy,
			"wheels":                       asset.Wheels,
			"model":                        asset.Model,
			"color":                        asset.Color,
			"fuel":                         asset.Fuel,
			"owning":                       asset.Owning,
			"brand":                        asset.Brand,
			"company":                      asset.Company,
			"company_id":                   asset.CompanyID,
			"type_name":                    asset.TypeName,
		})
	}
	c.JSON(http.StatusOK, response)
}

// Get returns a single vehicle asset
func (v *VehicleAssetController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// log.Printf("GET /v1/vehicle-asset/%d received\n", id)

	asset := models.VehicleAsset{ID: id}
	if err := asset.GetWithJoins(); err != nil {
		// log.Printf("Vehicle asset %d not found: %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	response := gin.H{
		"id":                           asset.ID,
		"license_plate":                asset.LicensePlate,
		"stnk_status":                  asset.STNKStatus,
		"bpkb_number":                  asset.BPKBNumber,
		"bpkb_status":                  asset.BPKBStatus,
		"chassis_number":               asset.ChassisNumber,
		"machine_number":               asset.MachineNumber,
		"cc_capacity":                  nullInt64ToJSON(asset.CCCapacity),
		"manufacture_year":             asset.ManufactureYear,
		"tax_due_date":                 nullInt64ToJSON(asset.TaxDueDate),
		"last_tax_payment_date":        nullInt64ToJSON(asset.LastTaxPaymentDate),
		"current_owner":                asset.CurrentOwner,
		"stnk_photo_path":              asset.STNKPhotoPath,
		"stnk_photo_verified":          asset.STNKPhotoVerified,
		"bpkb_photo_path":              asset.BPKBPhotoPath,
		"bpkb_photo_verified":          asset.BPKBPhotoVerified,
		"owner_id_photo_path":          asset.OwnerIDPhotoPath,
		"owner_id_photo_verified":      asset.OwnerIDPhotoVerified,
		"vehicle_photo_path":           asset.VehiclePhotoPath,
		"vehicle_photo_verified":       asset.VehiclePhotoVerified,
		"payment_billing_photo_path":   asset.PaymentBillingPhotoPath,
		"payment_billing_date":         nullInt64ToJSON(asset.PaymentBillingDate),
		"recommendation_document_path": asset.RecommendationDocPath,
		"e_sign_status":                asset.ESignStatus,
		"notes":                        asset.Notes,
		"status":                       asset.Status,
		"created_at":                   asset.CreatedAt,
		"updated_at":                   asset.UpdatedAt,
		"created_by":                   asset.CreatedBy,
		"wheels":                       asset.Wheels,
		"model":                        asset.Model,
		"color":                        asset.Color,
		"fuel":                         asset.Fuel,
		"owning":                       asset.Owning,
		"brand":                        asset.Brand,
		"company":                      asset.Company,
		"company_id":                   asset.CompanyID,
		"type_name":                    asset.TypeName,
	}
	c.JSON(http.StatusOK, response)
}

// Update modifies a vehicle asset (full update)
func (v *VehicleAssetController) Update(c *gin.Context) {
	var form forms.UpdateVehicleAssetForm
	// Handle JSON requests
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// First, get the existing asset to preserve photo paths and other fields
	asset := models.VehicleAsset{ID: int(form.ID)}
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found"})
		return
	}

	// Update only the fields that are provided in the form, preserving existing photo paths
	asset.LicensePlate = form.LicensePlate
	asset.STNKStatus = form.STNKStatus
	asset.BPKBNumber = form.BPKBNumber
	asset.BPKBStatus = form.BPKBStatus
	asset.ChassisNumber = form.ChassisNumber
	asset.MachineNumber = form.MachineNumber
	asset.TypeName = form.TypeName
	asset.WheelsID = sql.NullInt64{Int64: form.WheelsID.Int64, Valid: form.WheelsID.Valid}
	asset.ModelID = sql.NullInt64{Int64: form.ModelID.Int64, Valid: form.ModelID.Valid}
	asset.ColorID = sql.NullInt64{Int64: form.ColorID.Int64, Valid: form.ColorID.Valid}
	asset.FuelID = sql.NullInt64{Int64: form.FuelID.Int64, Valid: form.FuelID.Valid}
	asset.OwningID = sql.NullInt64{Int64: form.OwningID.Int64, Valid: form.OwningID.Valid}
	asset.BrandID = sql.NullInt64{Int64: form.BrandID.Int64, Valid: form.BrandID.Valid}
	asset.CCCapacity = sql.NullInt64{Int64: form.CCCapacity.Int64, Valid: form.CCCapacity.Valid}
	asset.ManufactureYear = int(form.ManufactureYear)
	asset.TaxDueDate = sql.NullInt64{Int64: form.TaxDueDate.Int64, Valid: form.TaxDueDate.Valid}
	asset.CurrentOwner = form.CurrentOwner
	asset.CompanyID = int(form.CompanyID.Int64)
	// Note: Status field is commented out in the form, so we don't update it

	if err := asset.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Delete removes a vehicle asset
func (v *VehicleAssetController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	asset := models.VehicleAsset{ID: id}
	if err := asset.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// GetByLicensePlate finds asset by plate
func (v *VehicleAssetController) GetByLicensePlate(c *gin.Context) {
	plate := c.Param("plate")
	asset := models.VehicleAsset{LicensePlate: plate}
	if err := asset.GetByLicensePlate(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, asset)
}

// GetByCompany returns all vehicle assets for a given company
func (v *VehicleAssetController) GetByCompany(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	asset := models.VehicleAsset{CompanyID: companyID}
	assets, err := asset.GetByCompany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format response for each asset
	var response []gin.H
	for _, asset := range assets {
		response = append(response, gin.H{
			"id":                           asset.ID,
			"license_plate":                asset.LicensePlate,
			"stnk_status":                  asset.STNKStatus,
			"bpkb_number":                  asset.BPKBNumber,
			"bpkb_status":                  asset.BPKBStatus,
			"chassis_number":               asset.ChassisNumber,
			"machine_number":               asset.MachineNumber,
			"cc_capacity":                  nullInt64ToJSON(asset.CCCapacity),
			"manufacture_year":             asset.ManufactureYear,
			"tax_due_date":                 nullInt64ToJSON(asset.TaxDueDate),
			"last_tax_payment_date":        nullInt64ToJSON(asset.LastTaxPaymentDate),
			"current_owner":                asset.CurrentOwner,
			"stnk_photo_path":              asset.STNKPhotoPath,
			"stnk_photo_verified":          asset.STNKPhotoVerified,
			"bpkb_photo_path":              asset.BPKBPhotoPath,
			"bpkb_photo_verified":          asset.BPKBPhotoVerified,
			"owner_id_photo_path":          asset.OwnerIDPhotoPath,
			"owner_id_photo_verified":      asset.OwnerIDPhotoVerified,
			"vehicle_photo_path":           asset.VehiclePhotoPath,
			"vehicle_photo_verified":       asset.VehiclePhotoVerified,
			"payment_billing_photo_path":   asset.PaymentBillingPhotoPath,
			"payment_billing_date":         nullInt64ToJSON(asset.PaymentBillingDate),
			"recommendation_document_path": asset.RecommendationDocPath,
			"e_sign_status":                asset.ESignStatus,
			"notes":                        asset.Notes,
			"status":                       asset.Status,
			"created_at":                   asset.CreatedAt,
			"updated_at":                   asset.UpdatedAt,
			"created_by":                   asset.CreatedBy,
			"wheels":                       asset.Wheels,
			"model":                        asset.Model,
			"color":                        asset.Color,
			"fuel":                         asset.Fuel,
			"owning":                       asset.Owning,
			"brand":                        asset.Brand,
			"company":                      asset.Company,
			"company_id":                   asset.CompanyID,
			"type_name":                    asset.TypeName,
		})
	}
	c.JSON(http.StatusOK, response)
}

// Reference Type APIs

func (v *VehicleAssetController) GetWheels(c *gin.Context) {
	wheels, err := models.GetAllVehicleWheels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wheels)
}

func (v *VehicleAssetController) GetModels(c *gin.Context) {
	models, err := models.GetAllVehicleModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}

func (v *VehicleAssetController) GetColors(c *gin.Context) {
	colors, err := models.GetAllVehicleColors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, colors)
}

func (v *VehicleAssetController) GetFuels(c *gin.Context) {
	fuels, err := models.GetAllVehicleFuels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fuels)
}

func (v *VehicleAssetController) GetCompanies(c *gin.Context) {
	companies, err := models.GetAllUtilizationCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (v *VehicleAssetController) GetOwningTypes(c *gin.Context) {
	owningTypes, err := models.GetAllVehicleOwningTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, owningTypes)
}

func (v *VehicleAssetController) GetBrands(c *gin.Context) {
	brands, err := models.GetAllVehicleBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

// VerifyDocument handles verifying/unverifying vehicle asset documents
func (v *VehicleAssetController) VerifyDocument(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle asset ID"})
		return
	}

	var form struct {
		Type     string `json:"type"`
		Verified bool   `json:"verified"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docType := form.Type
	if docType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Document type is required"})
		return
	}

	asset := models.VehicleAsset{ID: id}
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found", "details": err.Error()})
		return
	}

	// Update the appropriate verified field based on type
	switch docType {
	case "stnk":
		asset.STNKPhotoVerified = form.Verified
	case "bpkb":
		asset.BPKBPhotoVerified = form.Verified
	case "owner":
		asset.OwnerIDPhotoVerified = form.Verified
	case "vehicle":
		asset.VehiclePhotoVerified = form.Verified
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document type"})
		return
	}

	// Update status based on document verification (only STNK, owner, vehicle required)
	if asset.STNKPhotoPath != "" && asset.OwnerIDPhotoPath != "" && asset.VehiclePhotoPath != "" {
		if asset.STNKPhotoVerified && asset.OwnerIDPhotoVerified && asset.VehiclePhotoVerified {
			// All documents verified, check tax payment status
			currentYear := time.Now().Year()
			if asset.LastTaxPaymentDate.Valid {
				lastPaymentYear := time.Unix(int64(asset.LastTaxPaymentDate.Int64), 0).Year()
				if lastPaymentYear == currentYear {
					asset.Status = "paid"
				} else {
					// Check if RecommendationDocPath exists, if so set to unpaid instead of completed
					if asset.RecommendationDocPath != "" {
						asset.Status = "unpaid"
					} else {
						asset.Status = "completed"
					}
				}
			} else {
				// Check if RecommendationDocPath exists, if so set to unpaid instead of completed
				if asset.RecommendationDocPath != "" {
					asset.Status = "unpaid"
				} else {
					asset.Status = "completed"
				}
			}
		} else {
			asset.Status = "unverified"
		}
	} else {
		asset.Status = "uncompleted"
	}

	if err := asset.UpdateDocumentStatus(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle asset", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"message":      fmt.Sprintf("Document %s verification status updated", docType),
		"asset_status": asset.Status,
	})
}

// UploadDocument handles vehicle asset document uploads (STNK, BPKB, etc.)
func (v *VehicleAssetController) UploadDocument(c *gin.Context) {
	// Explicitly parse as multipart first to prevent JSON parsing
	contentType := c.GetHeader("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form", "details": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle asset ID"})
		return
	}

	docType := c.PostForm("type") // Changed from Query() to PostForm()
	if docType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Document type is required"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File upload is required"})
		return
	}
	defer file.Close()

	// Upload file to MinIO
	fileURL, err := services.UploadFile(header)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "File upload failed", "details": err.Error()})
		return
	}

	asset := models.VehicleAsset{ID: id}
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found"})
		return
	}

	// Update the appropriate document field based on type and set verified to false
	switch docType {
	case "stnk":
		asset.STNKPhotoPath = fileURL
		asset.STNKPhotoVerified = false
	case "bpkb":
		asset.BPKBPhotoPath = fileURL
		asset.BPKBPhotoVerified = false
	case "owner":
		asset.OwnerIDPhotoPath = fileURL
		asset.OwnerIDPhotoVerified = false
	case "vehicle":
		asset.VehiclePhotoPath = fileURL
		asset.VehiclePhotoVerified = false
	case "billing":
		asset.PaymentBillingPhotoPath = fileURL
	case "recommendation":
		asset.RecommendationDocPath = fileURL
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document type"})
		return
	}

	// Update status based on document completion - only for required documents (STNK, owner, vehicle)
	if docType == "stnk" || docType == "owner" || docType == "vehicle" {
		if asset.STNKPhotoPath != "" && asset.OwnerIDPhotoPath != "" && asset.VehiclePhotoPath != "" {
			asset.Status = "unverified"
		} else {
			asset.Status = "uncompleted"
		}
	}

	if err := asset.UpdateDocumentStatus(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update document paths", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"message":      "Document uploaded successfully",
		"url":          fileURL,
		"type":         docType,
		"asset_status": asset.Status,
	})
}

// RemoveDocument handles vehicle asset document removal (STNK, BPKB, etc.)
func (v *VehicleAssetController) RemoveDocument(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle asset ID"})
		return
	}

	var form struct {
		Type string `json:"type"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docType := form.Type
	if docType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Document type is required"})
		return
	}

	asset := models.VehicleAsset{ID: id}
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found"})
		return
	}

	var filePath string

	// Get the file path based on document type
	switch docType {
	case "stnk":
		filePath = asset.STNKPhotoPath
	case "bpkb":
		filePath = asset.BPKBPhotoPath
	case "owner":
		filePath = asset.OwnerIDPhotoPath
	case "vehicle":
		filePath = asset.VehiclePhotoPath
	case "billing":
		filePath = asset.PaymentBillingPhotoPath
	case "recommendation":
		filePath = asset.RecommendationDocPath
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document type"})
		return
	}

	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No document found to remove"})
		return
	}

	// Delete file from MinIO (only if it exists)
	// Note: We don't return error if file deletion fails, as the database record should still be cleared
	if filePath != "" {
		if err := services.DeleteObject(filePath); err != nil {
			// Log the error but continue with database cleanup
			log.Printf("Warning: Failed to delete file from storage: %v. Continuing with database cleanup.\n", err)
		}
	}

	// Special handling for billing document removal
	if docType == "billing" {
		// Clear the billing document path and date
		asset.PaymentBillingPhotoPath = ""
		asset.PaymentBillingDate = sql.NullInt64{Valid: false}

		// Check if all required documents are verified
		allRequiredVerified := asset.STNKPhotoVerified && asset.OwnerIDPhotoVerified && asset.VehiclePhotoVerified

		if allRequiredVerified {
			// If all verified, revert status to unpaid
			asset.Status = "unpaid"
		} else {
			// Check if all required documents exist (not necessarily verified)
			allRequiredExist := asset.STNKPhotoPath != "" && asset.OwnerIDPhotoPath != "" && asset.VehiclePhotoPath != ""

			if allRequiredExist {
				// If all required documents exist but not all verified, set to unverified
				asset.Status = "unverified"
			} else {
				// If not all required documents exist, set to uncompleted
				asset.Status = "uncompleted"
			}
		}

		// Revert tax_due_date to be last_payment_date
		asset.TaxDueDate = asset.LastTaxPaymentDate

		// Set last_payment_date to null
		asset.LastTaxPaymentDate = sql.NullInt64{Valid: false}

		// payment_billing_date is already set to null above

		// Use Update() method instead of UpdateDocumentStatus() since we're updating more fields
		if err := asset.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle asset", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":                "success",
			"message":               fmt.Sprintf("Document %s removed successfully", docType),
			"asset_status":          asset.Status,
			"tax_due_date":          nullInt64ToJSON(asset.TaxDueDate),
			"last_tax_payment_date": nullInt64ToJSON(asset.LastTaxPaymentDate),
			"payment_billing_date":  nullInt64ToJSON(asset.PaymentBillingDate),
		})
		return
	}

	// Clear the document path and verification status in database for non-billing documents
	switch docType {
	case "stnk":
		asset.STNKPhotoPath = ""
		asset.STNKPhotoVerified = false
	case "bpkb":
		asset.BPKBPhotoPath = ""
		asset.BPKBPhotoVerified = false
	case "owner":
		asset.OwnerIDPhotoPath = ""
		asset.OwnerIDPhotoVerified = false
	case "vehicle":
		asset.VehiclePhotoPath = ""
		asset.VehiclePhotoVerified = false
	case "recommendation":
		asset.RecommendationDocPath = ""
	}

	// Update status based on document completion - only for required documents (STNK, owner, vehicle)
	if docType == "stnk" || docType == "owner" || docType == "vehicle" {
		if asset.STNKPhotoPath != "" && asset.OwnerIDPhotoPath != "" && asset.VehiclePhotoPath != "" {
			asset.Status = "unverified"
		} else {
			asset.Status = "uncompleted"
		}
	}

	if err := asset.UpdateDocumentStatus(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle asset", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"message":      fmt.Sprintf("Document %s removed successfully", docType),
		"asset_status": asset.Status,
	})
}

// UpdatePayment handles payment update for vehicle asset
func (v *VehicleAssetController) UpdatePayment(c *gin.Context) {
	var form forms.UpdatePaymentForm

	// Handle JSON requests
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset := models.VehicleAsset{ID: int(form.ID)}

	// First, get the asset to check current state
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found"})
		return
	}

	// Use the UpdatePayment method which handles validation and updating
	if err := asset.UpdatePayment(int(form.TaxDueDate)); err != nil {
		// Handle specific validation errors
		if err.Error() == "Belum Upload Bukti Pembayaran" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "Belum Upload STNK" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Handle other database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment", "details": err.Error()})
		return
	}

	// Return success response with updated data
	c.JSON(http.StatusOK, gin.H{
		"status":                "success",
		"message":               "Payment updated successfully",
		"last_tax_payment_date": asset.LastTaxPaymentDate,
		"tax_due_date":          asset.TaxDueDate,
		"payment_billing_date":  asset.PaymentBillingDate,
		"payment":               asset.Status,
	})
}

// UpdateVehicleRecommendation handles updating vehicle asset recommendation and status fields
func (v *VehicleAssetController) UpdateVehicleRecommendation(c *gin.Context) {
	var form forms.UpdateVehicleRecommendationForm

	// Handle JSON requests
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset := models.VehicleAsset{ID: form.ID}

	// First, get the asset to check current state
	if err := asset.Get(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle asset not found"})
		return
	}

	// Update the fields from the form
	asset.RecommendationDocPath = string(form.RecommendationDocumentPath)
	asset.ESignStatus = form.ESignStatus
	asset.Notes = form.Notes
	asset.Status = form.Status
	asset.STNKStatus = form.STNKStatus
	asset.BPKBStatus = form.BPKBStatus

	// Use the UpdateRecommendation method
	if err := asset.UpdateRecommendation(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle recommendation", "details": err.Error()})
		return
	}

	// Return success response with updated data
	c.JSON(http.StatusOK, gin.H{
		"status":                       "success",
		"message":                      "Vehicle recommendation updated successfully",
		"recommendation_document_path": asset.RecommendationDocPath,
		"e_sign_status":                asset.ESignStatus,
		"notes":                        asset.Notes,
		"asset_status":                 asset.Status,
		"stnk_status":                  asset.STNKStatus,
		"bpkb_status":                  asset.BPKBStatus,
		"updated_at":                   asset.UpdatedAt,
	})
}

// GetSummary returns summary data grouped by company with overall summary
func (v *VehicleAssetController) GetSummary(c *gin.Context) {
	companyData, summary, err := models.GetExecutiveView()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"companies": companyData,
		"summary":   summary,
	}

	c.JSON(http.StatusOK, response)
}

// GetCompanySummary returns summary data for a specific company
func (v *VehicleAssetController) GetCompanySummary(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	companyData, err := models.GetCompanySummary(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companyData)
}

// UpdateTaxStatuses triggers the automated update of vehicle asset statuses based on tax due dates
func (v *VehicleAssetController) UpdateTaxStatuses(c *gin.Context) {
	err := models.UpdateTaxStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Tax statuses updated successfully"})
}
