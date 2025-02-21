package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// LkeKomponenForm ...
type LkeKomponenForm struct{}

// CreateLkeKomponenForm ...
type CreateLkeKomponenForm struct {
	KodeEvaluasi string  `form:"kode_evaluasi" json:"kode_evaluasi" binding:"required,min=3,max=50"`
	Bobot        float64 `form:"bobot" json:"bobot" binding:"required"`
	Komponen     string  `form:"komponen" json:"komponen" binding:"required,min=3,max=255"`
	Eviden       string  `form:"eviden" json:"eviden" binding:"required,min=3,max=255"`
	Level        string  `form:"level" json:"level" binding:"required,min=3,max=50"`
}

// KodeEvaluasi ...
func (f LkeKomponenForm) KodeEvaluasi(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the kode evaluasi"
		}
		return errMsg[0]
	case "min", "max":
		return "Kode evaluasi should be between 3 to 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Bobot ...
func (f LkeKomponenForm) Bobot(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the bobot value"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Komponen ...
func (f LkeKomponenForm) Komponen(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the komponen"
		}
		return errMsg[0]
	case "min", "max":
		return "Komponen should be between 3 to 255 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Eviden ...
func (f LkeKomponenForm) Eviden(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the eviden"
		}
		return errMsg[0]
	case "min", "max":
		return "Eviden should be between 3 to 255 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Level ...
func (f LkeKomponenForm) Level(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the level"
		}
		return errMsg[0]
	case "min", "max":
		return "Level should be between 3 to 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Create ...
func (f LkeKomponenForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "KodeEvaluasi" {
				return f.KodeEvaluasi(err.Tag())
			}
			if err.Field() == "Bobot" {
				return f.Bobot(err.Tag())
			}
			if err.Field() == "Komponen" {
				return f.Komponen(err.Tag())
			}
			if err.Field() == "Eviden" {
				return f.Eviden(err.Tag())
			}
			if err.Field() == "Level" {
				return f.Level(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f LkeKomponenForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "KodeEvaluasi" {
				return f.KodeEvaluasi(err.Tag())
			}
			if err.Field() == "Bobot" {
				return f.Bobot(err.Tag())
			}
			if err.Field() == "Komponen" {
				return f.Komponen(err.Tag())
			}
			if err.Field() == "Eviden" {
				return f.Eviden(err.Tag())
			}
			if err.Field() == "Level" {
				return f.Level(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
