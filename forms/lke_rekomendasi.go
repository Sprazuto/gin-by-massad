package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// LkeRekomendasiForm ...
type LkeRekomendasiForm struct{}

// CreateLkeRekomendasiForm ...
type CreateLkeRekomendasiForm struct {
	ParentID int64  `form:"parent_id" json:"parent_id" binding:"required"`
	Ta1a     string `form:"ta1a" json:"ta1a"`
	Ta1b     string `form:"ta1b" json:"ta1b"`
	Ta1c     string `form:"ta1c" json:"ta1c"`
	Ta2a     string `form:"ta2a" json:"ta2a"`
	Ta2b     string `form:"ta2b" json:"ta2b"`
	Ta2c     string `form:"ta2c" json:"ta2c"`
	Tb1      string `form:"tb1" json:"tb1"`
	Tb2      string `form:"tb2" json:"tb2"`
	Tb3      string `form:"tb3" json:"tb3"`
	Tc1      string `form:"tc1" json:"tc1"`
	Tc2      string `form:"tc2" json:"tc2"`
	Tc3      string `form:"tc3" json:"tc3"`
	Td1      string `form:"td1" json:"td1"`
	Td2      string `form:"td2" json:"td2"`
	Td3      string `form:"td3" json:"td3"`
}

// Create ...
func (f LkeRekomendasiForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "ParentID" {
				return "Please enter the parent ID"
			}
			if err.Field() == "Ta1a" {
				return "Please enter ta1a"
			}
			if err.Field() == "Ta1b" {
				return "Please enter ta1b"
			}
			if err.Field() == "Ta1c" {
				return "Please enter ta1c"
			}
			if err.Field() == "Ta2a" {
				return "Please enter ta2a"
			}
			if err.Field() == "Ta2b" {
				return "Please enter ta2b"
			}
			if err.Field() == "Ta2c" {
				return "Please enter ta2c"
			}
			if err.Field() == "Tb1" {
				return "Please enter tb1"
			}
			if err.Field() == "Tb2" {
				return "Please enter tb2"
			}
			if err.Field() == "Tb3" {
				return "Please enter tb3"
			}
			if err.Field() == "Tc1" {
				return "Please enter tc1"
			}
			if err.Field() == "Tc2" {
				return "Please enter tc2"
			}
			if err.Field() == "Tc3" {
				return "Please enter tc3"
			}
			if err.Field() == "Td1" {
				return "Please enter td1"
			}
			if err.Field() == "Td2" {
				return "Please enter td2"
			}
			if err.Field() == "Td3" {
				return "Please enter td3"
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
