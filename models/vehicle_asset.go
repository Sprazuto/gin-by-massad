package models

import (
	"database/sql"
	"fmt"
	"time"

	"aset-app/db"
)

// VehicleType represents a vehicle type reference
type VehicleType struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

// VehicleWheels represents a wheels configuration reference
type VehicleWheels struct {
	ID          int    `db:"id" json:"id"`
	Count       int    `db:"count" json:"count"`
	Description string `db:"description" json:"description"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

// VehicleModel represents a vehicle model reference
type VehicleModel struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

// VehicleColor represents a vehicle color reference
type VehicleColor struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	HexCode   string `db:"hex_code" json:"hex_code"`
	CreatedAt int    `db:"created_at" json:"created_at"`
	UpdatedAt int    `db:"updated_at" json:"updated_at"`
}

// VehicleFuel represents a fuel type reference
type VehicleFuel struct {
	ID          int    `db:"id" json:"id"`
	Type        string `db:"type" json:"type"`
	Description string `db:"description" json:"description"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

// UtilizationCompany represents a utilization company reference
type UtilizationCompany struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Address   string `db:"address" json:"address"`
	Contact   string `db:"contact" json:"contact"`
	CreatedAt int    `db:"created_at" json:"created_at"`
	UpdatedAt int    `db:"updated_at" json:"updated_at"`
}

// VehicleOwning represents a vehicle owning type reference
type VehicleOwning struct {
	ID          int    `db:"id" json:"id"`
	Type        string `db:"type" json:"type"`
	Description string `db:"description" json:"description"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

// VehicleBrand represents a vehicle brand reference
type VehicleBrand struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Country   string `db:"country" json:"country"`
	Founded   int    `db:"founded" json:"founded"`
	CreatedAt int    `db:"created_at" json:"created_at"`
	UpdatedAt int    `db:"updated_at" json:"updated_at"`
}

// VehicleAsset represents the vehicle asset model
type VehicleAsset struct {
	ID                      int           `db:"id" json:"id"`
	LicensePlate            string        `db:"license_plate" json:"license_plate"`
	STNKStatus              string        `db:"stnk_status" json:"stnk_status"`
	BPKBNumber              string        `db:"bpkb_number" json:"bpkb_number"`
	BPKBStatus              string        `db:"bpkb_status" json:"bpkb_status"`
	ChassisNumber           string        `db:"chassis_number" json:"chassis_number"`
	MachineNumber           string        `db:"machine_number" json:"machine_number"`
	TypeName                string        `db:"type_name" json:"type_name"`
	WheelsID                sql.NullInt64 `db:"wheels_id" json:"wheels_id"`
	ModelID                 sql.NullInt64 `db:"model_id" json:"model_id"`
	ColorID                 sql.NullInt64 `db:"color_id" json:"color_id"`
	FuelID                  sql.NullInt64 `db:"fuel_id" json:"fuel_id"`
	OwningID                sql.NullInt64 `db:"owning_id" json:"owning_id"`
	BrandID                 sql.NullInt64 `db:"brand_id" json:"brand_id"`
	CCCapacity              sql.NullInt64 `db:"cc_capacity" json:"cc_capacity"`
	ManufactureYear         int           `db:"manufacture_year" json:"manufacture_year"`
	TaxDueDate              sql.NullInt64 `db:"tax_due_date" json:"tax_due_date"`
	LastTaxPaymentDate      sql.NullInt64 `db:"last_tax_payment_date" json:"last_tax_payment_date"`
	CurrentOwner            string        `db:"current_owner" json:"current_owner"`
	CompanyID               int           `db:"company_id" json:"-"` // Hidden from JSON, will use custom field below
	APICompanyID            int           `db:"-" json:"company_id"` // Always include company_id in JSON response
	STNKPhotoPath           string        `db:"stnk_photo_path" json:"stnk_photo_path"`
	STNKPhotoVerified       bool          `db:"stnk_photo_verified" json:"stnk_photo_verified"`
	BPKBPhotoPath           string        `db:"bpkb_photo_path" json:"bpkb_photo_path"`
	BPKBPhotoVerified       bool          `db:"bpkb_photo_verified" json:"bpkb_photo_verified"`
	OwnerIDPhotoPath        string        `db:"owner_id_photo_path" json:"owner_id_photo_path"`
	OwnerIDPhotoVerified    bool          `db:"owner_id_photo_verified" json:"owner_id_photo_verified"`
	VehiclePhotoPath        string        `db:"vehicle_photo_path" json:"vehicle_photo_path"`
	VehiclePhotoVerified    bool          `db:"vehicle_photo_verified" json:"vehicle_photo_verified"`
	PaymentBillingPhotoPath string        `db:"payment_billing_photo_path" json:"payment_billing_photo_path"`
	PaymentBillingDate      sql.NullInt64 `db:"payment_billing_date" json:"payment_billing_date"`
	RecommendationDocPath   string        `db:"recommendation_document_path" json:"recommendation_document_path"`
	ESignStatus             bool          `db:"e_sign_status" json:"e_sign_status"`
	Notes                   string        `db:"notes" json:"notes"`
	Status                  string        `db:"status" json:"status"`
	CreatedAt               int           `db:"created_at" json:"created_at"`
	UpdatedAt               int           `db:"updated_at" json:"updated_at"`
	CreatedBy               int           `db:"created_by" json:"created_by"`

	// Nested reference objects
	Wheels  *VehicleWheels      `db:"-" json:"wheels,omitempty"`
	Model   *VehicleModel       `db:"-" json:"model,omitempty"`
	Color   *VehicleColor       `db:"-" json:"color,omitempty"`
	Fuel    *VehicleFuel        `db:"-" json:"fuel,omitempty"`
	Owning  *VehicleOwning      `db:"-" json:"owning,omitempty"`
	Brand   *VehicleBrand       `db:"-" json:"brand,omitempty"`
	Company *UtilizationCompany `db:"-" json:"company,omitempty"`
}

// Create creates a new vehicle asset
func (v *VehicleAsset) Create() (int, error) {
	now := int(time.Now().Unix())
	v.CreatedAt = now
	v.UpdatedAt = now

	query := `INSERT INTO vehicle_asset (
		license_plate, stnk_status, bpkb_number, bpkb_status, chassis_number, machine_number,
		type_name, wheels_id, model_id, color_id, fuel_id, owning_id, brand_id, cc_capacity,
		manufacture_year, tax_due_date, last_tax_payment_date, current_owner, company_id,
		notes, status, created_at, updated_at, created_by
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7,
		CASE WHEN $8 = 0 THEN NULL ELSE $8 END,
		CASE WHEN $9 = 0 THEN NULL ELSE $9 END,
		CASE WHEN $10 = 0 THEN NULL ELSE $10 END,
		CASE WHEN $11 = 0 THEN NULL ELSE $11 END,
		CASE WHEN $12 = 0 THEN NULL ELSE $12 END,
		CASE WHEN $13 = 0 THEN NULL ELSE $13 END,
		$14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24
	) RETURNING id`

	rows, err := db.GetDB().Query(query,
		v.LicensePlate, v.STNKStatus, v.BPKBNumber, v.BPKBStatus,
		v.ChassisNumber, v.MachineNumber, v.TypeName,
		v.WheelsID.Int64,
		v.ModelID.Int64,
		v.ColorID.Int64,
		v.FuelID.Int64,
		v.OwningID.Int64,
		v.BrandID.Int64,
		v.CCCapacity.Int64,
		v.ManufactureYear, v.TaxDueDate, v.LastTaxPaymentDate, v.CurrentOwner, v.CompanyID,
		v.Notes, v.Status, v.CreatedAt, v.UpdatedAt, v.CreatedBy)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	v.ID = id

	return id, nil
}

// Update updates a vehicle asset
func (v *VehicleAsset) Update() error {
	now := int(time.Now().Unix())
	v.UpdatedAt = now

	query := `UPDATE vehicle_asset SET
		license_plate = $1,
		stnk_status = $2,
		bpkb_number = $3,
		bpkb_status = $4,
		chassis_number = $5,
		machine_number = $6,
		type_name = $7,
		wheels_id = $8,
		model_id = $9,
		color_id = $10,
		fuel_id = $11,
		owning_id = $12,
		brand_id = $13,
		cc_capacity = $14,
		manufacture_year = $15,
		tax_due_date = $16,
		last_tax_payment_date = $17,
		current_owner = $18,
		company_id = $19,
		stnk_photo_path = $20,
		bpkb_photo_path = $21,
		owner_id_photo_path = $22,
		vehicle_photo_path = $23,
		payment_billing_photo_path = $24,
		recommendation_document_path = $25,
		e_sign_status = $26,
		notes = $27,
		status = $28,
		updated_at = $29,
		stnk_photo_verified = $30,
		bpkb_photo_verified = $31,
		owner_id_photo_verified = $32,
		vehicle_photo_verified = $33,
		payment_billing_date = $34
	WHERE id = $35`

	_, err := db.GetDB().Exec(query,
		v.LicensePlate, v.STNKStatus, v.BPKBNumber, v.BPKBStatus, v.ChassisNumber, v.MachineNumber,
		v.TypeName, v.WheelsID, v.ModelID, v.ColorID, v.FuelID, v.OwningID, v.BrandID,
		v.CCCapacity, v.ManufactureYear, v.TaxDueDate, v.LastTaxPaymentDate, v.CurrentOwner, v.CompanyID,
		v.STNKPhotoPath, v.BPKBPhotoPath, v.OwnerIDPhotoPath, v.VehiclePhotoPath,
		v.PaymentBillingPhotoPath, v.RecommendationDocPath, v.ESignStatus,
		v.Notes, v.Status, v.UpdatedAt,
		v.STNKPhotoVerified, v.BPKBPhotoVerified, v.OwnerIDPhotoVerified, v.VehiclePhotoVerified,
		v.PaymentBillingDate, v.ID)
	return err
}

// Get retrieves a vehicle asset by ID
func (v *VehicleAsset) Get() error {
	query := `
		SELECT
			id, license_plate,
			COALESCE(stnk_status, '') AS stnk_status,
			COALESCE(bpkb_number, '') AS bpkb_number,
			COALESCE(bpkb_status, '') AS bpkb_status,
			COALESCE(chassis_number, '') AS chassis_number,
			COALESCE(machine_number, '') AS machine_number,
			type_name, wheels_id, model_id, color_id, fuel_id, owning_id, brand_id, cc_capacity,
			manufacture_year, tax_due_date, last_tax_payment_date,
			COALESCE(current_owner, '') AS current_owner,
			company_id,
			COALESCE(stnk_photo_path, '') AS stnk_photo_path,
			stnk_photo_verified,
			COALESCE(bpkb_photo_path, '') AS bpkb_photo_path,
			bpkb_photo_verified,
			COALESCE(owner_id_photo_path, '') AS owner_id_photo_path,
			owner_id_photo_verified,
			COALESCE(vehicle_photo_path, '') AS vehicle_photo_path,
			vehicle_photo_verified,
			COALESCE(payment_billing_photo_path, '') AS payment_billing_photo_path,
			payment_billing_date,
			COALESCE(recommendation_document_path, '') AS recommendation_document_path,
			e_sign_status,
			COALESCE(notes, '') AS notes,
			COALESCE(status, '') AS status,
			created_at, updated_at, created_by
		FROM vehicle_asset WHERE id=$1`

	err := db.GetDB().SelectOne(v, query, v.ID)

	// Ensure company_id is set for JSON response
	v.APICompanyID = v.CompanyID

	return err
}

// GetAllVehicleAssets retrieves all vehicle assets with nested reference data
func GetAllVehicleAssets() ([]VehicleAsset, error) {
	// First get the base assets
	var baseAssets []VehicleAsset
	baseQuery := `SELECT
		id, license_plate,
		COALESCE(stnk_status, '') AS stnk_status,
		COALESCE(bpkb_number, '') AS bpkb_number,
		COALESCE(bpkb_status, '') AS bpkb_status,
		COALESCE(chassis_number, '') AS chassis_number,
		COALESCE(machine_number, '') AS machine_number,
		type_name,
		wheels_id,
		model_id,
		color_id,
		fuel_id,
		owning_id,
		brand_id,
		cc_capacity,
		manufacture_year, tax_due_date, last_tax_payment_date,
		COALESCE(current_owner, '') AS current_owner,
		company_id,
		COALESCE(stnk_photo_path, '') AS stnk_photo_path,
		stnk_photo_verified,
		COALESCE(bpkb_photo_path, '') AS bpkb_photo_path,
		bpkb_photo_verified,
		COALESCE(owner_id_photo_path, '') AS owner_id_photo_path,
		owner_id_photo_verified,
		COALESCE(vehicle_photo_path, '') AS vehicle_photo_path,
		vehicle_photo_verified,
		COALESCE(payment_billing_photo_path, '') AS payment_billing_photo_path,
		payment_billing_date,
		COALESCE(recommendation_document_path, '') AS recommendation_document_path,
		e_sign_status,
		COALESCE(notes, '') AS notes,
		COALESCE(status, '') AS status,
		created_at, updated_at, created_by
		FROM vehicle_asset ORDER BY id DESC`

	_, err := db.GetDB().Select(&baseAssets, baseQuery)
	if err != nil {
		return nil, err
	}

	// For each asset, get the nested objects if they exist
	var fullAssets []VehicleAsset
	for _, asset := range baseAssets {
		// Get related objects only if the foreign key is valid
		if asset.WheelsID.Valid {
			wheelsData := VehicleWheels{ID: int(asset.WheelsID.Int64)}
			if err := db.GetDB().SelectOne(&wheelsData, "SELECT * FROM vehicle_wheels WHERE id=$1", asset.WheelsID.Int64); err == nil {
				asset.Wheels = &wheelsData
			}
		}

		if asset.ModelID.Valid {
			modelData := VehicleModel{ID: int(asset.ModelID.Int64)}
			if err := db.GetDB().SelectOne(&modelData, "SELECT * FROM vehicle_model WHERE id=$1", asset.ModelID.Int64); err == nil {
				asset.Model = &modelData
			}
		}

		if asset.ColorID.Valid {
			colorData := VehicleColor{ID: int(asset.ColorID.Int64)}
			if err := db.GetDB().SelectOne(&colorData, "SELECT * FROM vehicle_color WHERE id=$1", asset.ColorID.Int64); err == nil {
				asset.Color = &colorData
			}
		}

		if asset.FuelID.Valid {
			fuelData := VehicleFuel{ID: int(asset.FuelID.Int64)}
			if err := db.GetDB().SelectOne(&fuelData, "SELECT * FROM vehicle_fuel WHERE id=$1", asset.FuelID.Int64); err == nil {
				asset.Fuel = &fuelData
			}
		}

		if asset.OwningID.Valid {
			owningData := VehicleOwning{ID: int(asset.OwningID.Int64)}
			if err := db.GetDB().SelectOne(&owningData, "SELECT * FROM vehicle_owning WHERE id=$1", asset.OwningID.Int64); err == nil {
				asset.Owning = &owningData
			}
		}

		if asset.BrandID.Valid {
			brandData := VehicleBrand{ID: int(asset.BrandID.Int64)}
			if err := db.GetDB().SelectOne(&brandData, "SELECT * FROM vehicle_brand WHERE id=$1", asset.BrandID.Int64); err == nil {
				asset.Brand = &brandData
			}
		}

		// Try to get company data if it exists locally (API integration flexibility)
		companyData := UtilizationCompany{ID: asset.CompanyID}
		if err := db.GetDB().SelectOne(&companyData, "SELECT * FROM utilization_company WHERE id=$1", asset.CompanyID); err == nil {
			asset.Company = &companyData
		}
		// If company doesn't exist locally, Company remains nil - this is expected for API integration

		// Ensure company_id is set for JSON response
		asset.APICompanyID = asset.CompanyID

		fullAssets = append(fullAssets, asset)
	}

	return fullAssets, nil
}

// Delete removes a vehicle asset
func (v *VehicleAsset) Delete() error {
	_, err := db.GetDB().Exec("DELETE FROM vehicle_asset WHERE id=$1", v.ID)
	return err
}

// UpdateDocument updates only document-related fields
func (v *VehicleAsset) UpdateDocument() error {
	now := int(time.Now().Unix())
	v.UpdatedAt = now

	query := `UPDATE vehicle_asset SET
		stnk_photo_path = $1,
		bpkb_photo_path = $2,
		owner_id_photo_path = $3,
		vehicle_photo_path = $4,
		payment_billing_photo_path = $5,
		recommendation_document_path = $6,
		updated_at = $7
	WHERE id = $8`

	_, err := db.GetDB().Exec(query,
		v.STNKPhotoPath, v.BPKBPhotoPath, v.OwnerIDPhotoPath, v.VehiclePhotoPath,
		v.PaymentBillingPhotoPath, v.RecommendationDocPath,
		v.UpdatedAt, v.ID)
	return err
}

// UpdateDocumentStatus updates document paths, verification status, and asset status without affecting other fields
func (v *VehicleAsset) UpdateDocumentStatus() error {
	now := int(time.Now().Unix())
	v.UpdatedAt = now

	query := `UPDATE vehicle_asset SET
		stnk_photo_path = $1,
		bpkb_photo_path = $2,
		owner_id_photo_path = $3,
		vehicle_photo_path = $4,
		payment_billing_photo_path = $5,
		recommendation_document_path = $6,
		stnk_photo_verified = $7,
		bpkb_photo_verified = $8,
		owner_id_photo_verified = $9,
		vehicle_photo_verified = $10,
		status = $11,
		updated_at = $12
	WHERE id = $13`

	_, err := db.GetDB().Exec(query,
		v.STNKPhotoPath, v.BPKBPhotoPath, v.OwnerIDPhotoPath, v.VehiclePhotoPath,
		v.PaymentBillingPhotoPath, v.RecommendationDocPath,
		v.STNKPhotoVerified, v.BPKBPhotoVerified, v.OwnerIDPhotoVerified, v.VehiclePhotoVerified,
		v.Status, v.UpdatedAt, v.ID)
	return err
}

// UpdatePayment updates payment-related fields with validation
func (v *VehicleAsset) UpdatePayment(newTaxDueDate int) error {
	// Check if required photo paths exist (only STNK, owner, vehicle required)
	if v.PaymentBillingPhotoPath == "" {
		return fmt.Errorf("Belum Upload Bukti Pembayaran")
	}
	if v.STNKPhotoPath == "" {
		return fmt.Errorf("Belum Upload STNK")
	}
	// Note: BPKB is no longer required for payment validation

	now := int(time.Now().Unix())
	v.UpdatedAt = now

	// Set old tax_due_date to last_tax_payment_date
	oldTaxDueDate := v.TaxDueDate

	// Create new tax due date as sql.NullInt64
	newTaxDueDateNull := sql.NullInt64{Int64: int64(newTaxDueDate), Valid: true}

	// Create payment billing date as current timestamp
	paymentBillingDate := sql.NullInt64{Int64: int64(now), Valid: true}

	// Update the payment-related fields
	query := `UPDATE vehicle_asset SET
		last_tax_payment_date = $1,
		tax_due_date = $2,
		payment_billing_date = $3,
		status = $4,
		updated_at = $5
	WHERE id = $6`

	_, err := db.GetDB().Exec(query,
		oldTaxDueDate, newTaxDueDateNull, paymentBillingDate, "paid", now, v.ID)

	if err == nil {
		// Update the model instance with new values
		v.LastTaxPaymentDate = oldTaxDueDate
		v.TaxDueDate = newTaxDueDateNull
		v.PaymentBillingDate = paymentBillingDate
		v.Status = "paid"
	}

	return err
}

// UpdateRecommendation updates recommendation and status-related fields
func (v *VehicleAsset) UpdateRecommendation() error {
	now := int(time.Now().Unix())
	v.UpdatedAt = now

	query := `UPDATE vehicle_asset SET
		recommendation_document_path = $1,
		e_sign_status = $2,
		notes = $3,
		status = $4,
		stnk_status = $5,
		bpkb_status = $6,
		updated_at = $7
	WHERE id = $8`

	_, err := db.GetDB().Exec(query,
		v.RecommendationDocPath, v.ESignStatus, v.Notes, v.Status,
		v.STNKStatus, v.BPKBStatus, now, v.ID)
	return err
}

// Reference Type Methods
func GetAllVehicleWheels() ([]VehicleWheels, error) {
	var wheels []VehicleWheels
	_, err := db.GetDB().Select(&wheels, "SELECT * FROM vehicle_wheels ORDER BY id")
	return wheels, err
}

func GetAllVehicleModels() ([]VehicleModel, error) {
	var models []VehicleModel
	_, err := db.GetDB().Select(&models, "SELECT * FROM vehicle_model ORDER BY id")
	return models, err
}

func GetAllVehicleColors() ([]VehicleColor, error) {
	var colors []VehicleColor
	_, err := db.GetDB().Select(&colors, "SELECT * FROM vehicle_color ORDER BY id")
	return colors, err
}

func GetAllVehicleFuels() ([]VehicleFuel, error) {
	var fuels []VehicleFuel
	_, err := db.GetDB().Select(&fuels, "SELECT * FROM vehicle_fuel ORDER BY id")
	return fuels, err
}

func GetAllUtilizationCompanies() ([]UtilizationCompany, error) {
	var companies []UtilizationCompany
	_, err := db.GetDB().Select(&companies, "SELECT * FROM utilization_company ORDER BY id")
	return companies, err
}

func GetAllVehicleOwningTypes() ([]VehicleOwning, error) {
	var types []VehicleOwning
	_, err := db.GetDB().Select(&types, "SELECT * FROM vehicle_owning ORDER BY id")
	return types, err
}

func GetAllVehicleBrands() ([]VehicleBrand, error) {
	var brands []VehicleBrand
	_, err := db.GetDB().Select(&brands, "SELECT * FROM vehicle_brand ORDER BY id")
	return brands, err
}

// GetByLicensePlate retrieves a vehicle asset by license plate
func (v *VehicleAsset) GetByLicensePlate() error {
	err := db.GetDB().SelectOne(v, "SELECT * FROM vehicle_asset WHERE license_plate=$1", v.LicensePlate)
	return err
}

// GetByCompany retrieves vehicle assets by company ID with nested reference data
func (v *VehicleAsset) GetByCompany() ([]VehicleAsset, error) {
	// First get the base assets
	var baseAssets []VehicleAsset
	baseQuery := `SELECT
		id, license_plate,
		COALESCE(stnk_status, '') AS stnk_status,
		COALESCE(bpkb_number, '') AS bpkb_number,
		COALESCE(bpkb_status, '') AS bpkb_status,
		COALESCE(chassis_number, '') AS chassis_number,
		COALESCE(machine_number, '') AS machine_number,
		type_name,
		wheels_id,
		model_id,
		color_id,
		fuel_id,
		owning_id,
		brand_id,
		cc_capacity,
		manufacture_year, tax_due_date, last_tax_payment_date,
		COALESCE(current_owner, '') AS current_owner,
		company_id,
		COALESCE(stnk_photo_path, '') AS stnk_photo_path,
		stnk_photo_verified,
		COALESCE(bpkb_photo_path, '') AS bpkb_photo_path,
		bpkb_photo_verified,
		COALESCE(owner_id_photo_path, '') AS owner_id_photo_path,
		owner_id_photo_verified,
		COALESCE(vehicle_photo_path, '') AS vehicle_photo_path,
		vehicle_photo_verified,
		COALESCE(payment_billing_photo_path, '') AS payment_billing_photo_path,
		payment_billing_date,
		COALESCE(recommendation_document_path, '') AS recommendation_document_path,
		e_sign_status,
		COALESCE(notes, '') AS notes,
		COALESCE(status, '') AS status,
		created_at, updated_at, created_by
		FROM vehicle_asset WHERE company_id=$1 ORDER BY id`

	_, err := db.GetDB().Select(&baseAssets, baseQuery, v.CompanyID)
	if err != nil {
		return nil, err
	}

	// For each asset, get the nested objects if they exist
	var fullAssets []VehicleAsset
	for _, asset := range baseAssets {
		// Get related objects only if the foreign key is valid
		if asset.WheelsID.Valid {
			wheelsData := VehicleWheels{ID: int(asset.WheelsID.Int64)}
			if err := db.GetDB().SelectOne(&wheelsData, "SELECT * FROM vehicle_wheels WHERE id=$1", asset.WheelsID.Int64); err == nil {
				asset.Wheels = &wheelsData
			}
		}

		if asset.ModelID.Valid {
			modelData := VehicleModel{ID: int(asset.ModelID.Int64)}
			if err := db.GetDB().SelectOne(&modelData, "SELECT * FROM vehicle_model WHERE id=$1", asset.ModelID.Int64); err == nil {
				asset.Model = &modelData
			}
		}

		if asset.ColorID.Valid {
			colorData := VehicleColor{ID: int(asset.ColorID.Int64)}
			if err := db.GetDB().SelectOne(&colorData, "SELECT * FROM vehicle_color WHERE id=$1", asset.ColorID.Int64); err == nil {
				asset.Color = &colorData
			}
		}

		if asset.FuelID.Valid {
			fuelData := VehicleFuel{ID: int(asset.FuelID.Int64)}
			if err := db.GetDB().SelectOne(&fuelData, "SELECT * FROM vehicle_fuel WHERE id=$1", asset.FuelID.Int64); err == nil {
				asset.Fuel = &fuelData
			}
		}

		if asset.OwningID.Valid {
			owningData := VehicleOwning{ID: int(asset.OwningID.Int64)}
			if err := db.GetDB().SelectOne(&owningData, "SELECT * FROM vehicle_owning WHERE id=$1", asset.OwningID.Int64); err == nil {
				asset.Owning = &owningData
			}
		}

		if asset.BrandID.Valid {
			brandData := VehicleBrand{ID: int(asset.BrandID.Int64)}
			if err := db.GetDB().SelectOne(&brandData, "SELECT * FROM vehicle_brand WHERE id=$1", asset.BrandID.Int64); err == nil {
				asset.Brand = &brandData
			}
		}

		// Try to get company data if it exists locally (API integration flexibility)
		companyData := UtilizationCompany{ID: asset.CompanyID}
		if err := db.GetDB().SelectOne(&companyData, "SELECT * FROM utilization_company WHERE id=$1", asset.CompanyID); err == nil {
			asset.Company = &companyData
		}
		// If company doesn't exist locally, Company remains nil - this is expected for API integration

		fullAssets = append(fullAssets, asset)
	}

	return fullAssets, nil
}

// GetWithJoins retrieves a vehicle asset with nested joined reference data
func (v *VehicleAsset) GetWithJoins() error {
	// First get the base vehicle asset
	err := v.Get()
	if err != nil {
		return err
	}

	if v.WheelsID.Valid {
		wheelsData := VehicleWheels{ID: int(v.WheelsID.Int64)}
		if err := db.GetDB().SelectOne(&wheelsData, "SELECT * FROM vehicle_wheels WHERE id=$1", v.WheelsID.Int64); err != nil {
			return err
		}
		v.Wheels = &wheelsData
	}

	if v.ModelID.Valid {
		modelData := VehicleModel{ID: int(v.ModelID.Int64)}
		if err := db.GetDB().SelectOne(&modelData, "SELECT * FROM vehicle_model WHERE id=$1", v.ModelID.Int64); err != nil {
			return err
		}
		v.Model = &modelData
	}

	if v.ColorID.Valid {
		colorData := VehicleColor{ID: int(v.ColorID.Int64)}
		if err := db.GetDB().SelectOne(&colorData, "SELECT * FROM vehicle_color WHERE id=$1", v.ColorID.Int64); err != nil {
			return err
		}
		v.Color = &colorData
	}

	if v.FuelID.Valid {
		fuelData := VehicleFuel{ID: int(v.FuelID.Int64)}
		if err := db.GetDB().SelectOne(&fuelData, "SELECT * FROM vehicle_fuel WHERE id=$1", v.FuelID.Int64); err != nil {
			return err
		}
		v.Fuel = &fuelData
	}

	if v.OwningID.Valid {
		owningData := VehicleOwning{ID: int(v.OwningID.Int64)}
		if err := db.GetDB().SelectOne(&owningData, "SELECT * FROM vehicle_owning WHERE id=$1", v.OwningID.Int64); err != nil {
			return err
		}
		v.Owning = &owningData
	}

	if v.BrandID.Valid {
		brandData := VehicleBrand{ID: int(v.BrandID.Int64)}
		if err := db.GetDB().SelectOne(&brandData, "SELECT * FROM vehicle_brand WHERE id=$1", v.BrandID.Int64); err != nil {
			return err
		}
		v.Brand = &brandData
	}

	// Try to get company data if it exists locally (API integration flexibility)
	companyData := UtilizationCompany{ID: v.CompanyID}
	if err := db.GetDB().SelectOne(&companyData, "SELECT * FROM utilization_company WHERE id=$1", v.CompanyID); err == nil {
		v.Company = &companyData
	}
	// If company doesn't exist locally, Company remains nil - this is acceptable for API integration

	return nil
}

// ExecutiveViewData represents the structure for executive view response
type ExecutiveViewData struct {
	CompanyID         int    `db:"company_id" json:"company_id"`
	CompanyName       string `db:"company_name" json:"company_name"`
	TotalAssets       int    `db:"total_assets" json:"total_assets"`
	UnreviewedAssets  int    `db:"unreviewed_assets" json:"unreviewed_assets"`
	UncompletedAssets int    `db:"uncompleted_assets" json:"uncompleted_assets"`
	UnverifiedAssets  int    `db:"unverified_assets" json:"unverified_assets"`
	CompletedAssets   int    `db:"completed_assets" json:"completed_assets"`
	UnsignedAssets    int    `db:"unsigned_assets" json:"unsigned_assets"`
	UnpaidAssets      int    `db:"unpaid_assets" json:"unpaid_assets"`
	PaidAssets        int    `db:"paid_assets" json:"paid_assets"`
}

// ExecutiveViewSummary represents the summary of all companies
type ExecutiveViewSummary struct {
	TotalCompanies    int `json:"total_companies"`
	TotalAssets       int `json:"total_assets"`
	UnreviewedAssets  int `json:"unreviewed_assets"`
	UncompletedAssets int `json:"uncompleted_assets"`
	UnverifiedAssets  int `json:"unverified_assets"`
	CompletedAssets   int `json:"completed_assets"`
	UnsignedAssets    int `json:"unsigned_assets"`
	UnpaidAssets      int `json:"unpaid_assets"`
	PaidAssets        int `json:"paid_assets"`
}

// GetExecutiveView retrieves executive view data grouped by company
func GetExecutiveView() ([]ExecutiveViewData, ExecutiveViewSummary, error) {
	// Query to get grouped data by company
	query := `
		SELECT
			va.company_id,
			COALESCE(uc.name, 'Unknown') as company_name,
			COUNT(va.id) as total_assets,
			COUNT(CASE WHEN va.status = 'unreviewed' THEN 1 END) as unreviewed_assets,
			COUNT(CASE WHEN va.status = 'uncompleted' THEN 1 END) as uncompleted_assets,
			COUNT(CASE WHEN va.status = 'unverified' THEN 1 END) as unverified_assets,
			COUNT(CASE WHEN va.status = 'completed' THEN 1 END) as completed_assets,
			COUNT(CASE WHEN va.status = 'unsigned' THEN 1 END) as unsigned_assets,
			COUNT(CASE WHEN va.status = 'unpaid' THEN 1 END) as unpaid_assets,
			COUNT(CASE WHEN va.status = 'paid' THEN 1 END) as paid_assets
		FROM vehicle_asset va
		LEFT JOIN utilization_company uc ON va.company_id = uc.id
		WHERE va.company_id IS NOT NULL AND va.company_id > 0
		GROUP BY va.company_id, uc.name
		ORDER BY uc.name NULLS LAST
	`

	rows, err := db.GetDB().Query(query)
	if err != nil {
		return nil, ExecutiveViewSummary{}, err
	}
	defer rows.Close()

	var companyData []ExecutiveViewData
	var summary ExecutiveViewSummary

	for rows.Next() {
		var data ExecutiveViewData
		err := rows.Scan(
			&data.CompanyID,
			&data.CompanyName,
			&data.TotalAssets,
			&data.UnreviewedAssets,
			&data.UncompletedAssets,
			&data.UnverifiedAssets,
			&data.CompletedAssets,
			&data.UnsignedAssets,
			&data.UnpaidAssets,
			&data.PaidAssets,
		)
		if err != nil {
			return nil, ExecutiveViewSummary{}, err
		}

		companyData = append(companyData, data)

		// Update summary
		summary.TotalCompanies++
		summary.TotalAssets += data.TotalAssets
		summary.UnreviewedAssets += data.UnreviewedAssets
		summary.UncompletedAssets += data.UncompletedAssets
		summary.UnverifiedAssets += data.UnverifiedAssets
		summary.CompletedAssets += data.CompletedAssets
		summary.UnsignedAssets += data.UnsignedAssets
		summary.UnpaidAssets += data.UnpaidAssets
		summary.PaidAssets += data.PaidAssets
	}

	if err = rows.Err(); err != nil {
		return nil, ExecutiveViewSummary{}, err
	}

	return companyData, summary, nil
}

// GetCompanySummary retrieves summary data for a specific company
func GetCompanySummary(companyID int) (ExecutiveViewData, error) {
	// Query to get summary data for a specific company
	query := `
		SELECT
			va.company_id,
			COALESCE(uc.name, 'Unknown') as company_name,
			COUNT(va.id) as total_assets,
			COUNT(CASE WHEN va.status = 'unreviewed' THEN 1 END) as unreviewed_assets,
			COUNT(CASE WHEN va.status = 'uncompleted' THEN 1 END) as uncompleted_assets,
			COUNT(CASE WHEN va.status = 'unverified' THEN 1 END) as unverified_assets,
			COUNT(CASE WHEN va.status = 'completed' THEN 1 END) as completed_assets,
			COUNT(CASE WHEN va.status = 'unsigned' THEN 1 END) as unsigned_assets,
			COUNT(CASE WHEN va.status = 'unpaid' THEN 1 END) as unpaid_assets,
			COUNT(CASE WHEN va.status = 'paid' THEN 1 END) as paid_assets
		FROM vehicle_asset va
		LEFT JOIN utilization_company uc ON va.company_id = uc.id
		WHERE va.company_id = $1
		GROUP BY va.company_id, uc.name
	`

	var data ExecutiveViewData
	err := db.GetDB().SelectOne(&data, query, companyID)
	if err != nil {
		return ExecutiveViewData{}, err
	}

	return data, nil
}
