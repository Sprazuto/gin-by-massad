package forms

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin/binding"
)

// NullInt64 is a custom type that implements json.Unmarshaler to handle string/int values
type NullInt64 sql.NullInt64

func (n *NullInt64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var num int64
	if err := json.Unmarshal(data, &num); err == nil {
		n.Int64 = num
		n.Valid = true
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == "" {
		n.Valid = false
		return nil
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	n.Int64 = num
	n.Valid = true
	return nil
}

// FlexInt handles string/int values for non-nullable integers
type FlexInt int

func (f *FlexInt) UnmarshalJSON(data []byte) error {
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FlexInt(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*f = FlexInt(num)
	return nil
}

// CreateVehicleAssetForm represents the vehicle asset create form
type CreateVehicleAssetForm struct {
	LicensePlate    string    `form:"license_plate" json:"license_plate" binding:"required"`
	STNKStatus      string    `form:"stnk_status" json:"stnk_status"`
	BPKBNumber      string    `form:"bpkb_number" json:"bpkb_number"`
	BPKBStatus      string    `form:"bpkb_status" json:"bpkb_status"`
	ChassisNumber   string    `form:"chassis_number" json:"chassis_number" binding:"required"`
	MachineNumber   string    `form:"machine_number" json:"machine_number" binding:"required"`
	TypeName        string    `form:"type_name" json:"type_name" binding:"required"`
	WheelsID        NullInt64 `form:"wheels_id" json:"wheels_id"`
	ModelID         NullInt64 `form:"model_id" json:"model_id"`
	ColorID         NullInt64 `form:"color_id" json:"color_id"`
	FuelID          NullInt64 `form:"fuel_id" json:"fuel_id"`
	OwningID        NullInt64 `form:"owning_id" json:"owning_id"`
	BrandID         NullInt64 `form:"brand_id" json:"brand_id"`
	CCCapacity      NullInt64 `form:"cc_capacity" json:"cc_capacity"`
	ManufactureYear FlexInt   `form:"manufacture_year" json:"manufacture_year" binding:"required,numeric"`
	TaxDueDate      NullInt64 `form:"tax_due_date" json:"tax_due_date"`
	CurrentOwner    string    `form:"current_owner" json:"current_owner" binding:"required"`
	CompanyID       NullInt64 `form:"company_id" json:"company_id" binding:"required"`
	Status          string    `form:"status" json:"status" binding:"required"`
}

// UpdateVehicleAssetForm represents the vehicle asset update form
type UpdateVehicleAssetForm struct {
	ID              FlexInt   `form:"id" json:"id" binding:"required"` // Changed to FlexInt to handle both string and int
	LicensePlate    string    `form:"license_plate" json:"license_plate" binding:"required"`
	STNKStatus      string    `form:"stnk_status" json:"stnk_status"`
	BPKBNumber      string    `form:"bpkb_number" json:"bpkb_number"`
	BPKBStatus      string    `form:"bpkb_status" json:"bpkb_status"`
	ChassisNumber   string    `form:"chassis_number" json:"chassis_number" binding:"required"`
	MachineNumber   string    `form:"machine_number" json:"machine_number" binding:"required"`
	TypeName        string    `form:"type_name" json:"type_name" binding:"required"`
	WheelsID        NullInt64 `form:"wheels_id" json:"wheels_id"`
	ModelID         NullInt64 `form:"model_id" json:"model_id"`
	ColorID         NullInt64 `form:"color_id" json:"color_id"`
	FuelID          NullInt64 `form:"fuel_id" json:"fuel_id"`
	OwningID        NullInt64 `form:"owning_id" json:"owning_id"`
	BrandID         NullInt64 `form:"brand_id" json:"brand_id"`
	CCCapacity      NullInt64 `form:"cc_capacity" json:"cc_capacity"`
	ManufactureYear FlexInt   `form:"manufacture_year" json:"manufacture_year" binding:"required,numeric"`
	TaxDueDate      NullInt64 `form:"tax_due_date" json:"tax_due_date"`
	CurrentOwner    string    `form:"current_owner" json:"current_owner" binding:"required"`
	CompanyID       NullInt64 `form:"company_id" json:"company_id" binding:"required"`
	// Status          string    `form:"status" json:"status" binding:"required"`
}

// Valid validates the CreateVehicleAssetForm
func (f *CreateVehicleAssetForm) Valid() error {
	return binding.Validator.ValidateStruct(f)
}

// Valid validates the UpdateVehicleAssetForm
func (f *UpdateVehicleAssetForm) Valid() error {
	return binding.Validator.ValidateStruct(f)
}

// UpdatePaymentForm represents the payment update form
type UpdatePaymentForm struct {
	ID         FlexInt `form:"id" json:"id" binding:"required"`
	TaxDueDate FlexInt `form:"tax_due_date" json:"tax_due_date" binding:"required"`
}

// Valid validates the UpdatePaymentForm
func (f *UpdatePaymentForm) Valid() error {
	return binding.Validator.ValidateStruct(f)
}

// FlexString handles string/number values for string fields
type FlexString string

func (f *FlexString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*f = FlexString(str)
		return nil
	}

	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FlexString(fmt.Sprintf("%.0f", num))
		return nil
	}

	return fmt.Errorf("invalid value for FlexString: %s", string(data))
}

// UpdateVehicleRecommendationForm represents the form for updating vehicle asset recommendation and status fields
type UpdateVehicleRecommendationForm struct {
	ID                         int        `form:"id" json:"id" binding:"required"`
	RecommendationDocumentPath FlexString `form:"recommendation_document_path" json:"recommendation_document_path"`
	ESignStatus                bool       `form:"e_sign_status" json:"e_sign_status"`
	Notes                      string     `form:"notes" json:"notes"`
	Status                     string     `form:"status" json:"status"`
	STNKStatus                 string     `form:"stnk_status" json:"stnk_status"`
	BPKBStatus                 string     `form:"bpkb_status" json:"bpkb_status"`
}

// Valid validates the UpdateVehicleRecommendationForm
func (f *UpdateVehicleRecommendationForm) Valid() error {
	return binding.Validator.ValidateStruct(f)
}

// Init initializes the form with
