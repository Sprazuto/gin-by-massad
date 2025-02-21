package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// LkeEvaluasiForm ...
type LkeEvaluasiForm struct{}

// CreateLkeEvaluasiForm ...
type CreateLkeEvaluasiForm struct {
	LkeRekapID   int64   `form:"lke_rekap_id" json:"lke_rekap_id" binding:"required"`
	KodeEvaluasi string  `form:"kode_evaluasi" json:"kode_evaluasi" binding:"required,max=50"`
	Jawaban      *string `form:"jawaban" json:"jawaban"`
	Berkas       *string `form:"berkas" json:"berkas"`
	Catatan      *string `form:"catatan" json:"catatan"`
}

// LkeRekapID ...
func (f LkeEvaluasiForm) LkeRekapID(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the LKE Rekap ID"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// KodeEvaluasi ...
func (f LkeEvaluasiForm) KodeEvaluasi(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the kode evaluasi"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Jawaban ...
func (f LkeEvaluasiForm) Jawaban(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the jawaban"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Berkas ...
func (f LkeEvaluasiForm) Berkas(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the berkas"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Catatan ...
func (f LkeEvaluasiForm) Catatan(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the catatan"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Create ...
func (f LkeEvaluasiForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "LkeRekapID":
				return f.LkeRekapID(err.Tag())
			case "KodeEvaluasi":
				return f.KodeEvaluasi(err.Tag())
			case "Jawaban":
				return f.Jawaban(err.Tag())
			case "Berkas":
				return f.Berkas(err.Tag())
			case "Catatan":
				return f.Catatan(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f LkeEvaluasiForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "LkeRekapID":
				return f.LkeRekapID(err.Tag())
			case "KodeEvaluasi":
				return f.KodeEvaluasi(err.Tag())
			case "Jawaban":
				return f.Jawaban(err.Tag())
			case "Berkas":
				return f.Berkas(err.Tag())
			case "Catatan":
				return f.Catatan(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
