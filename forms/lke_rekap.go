package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// LkeRekapForm ...
type LkeRekapForm struct{}

// CreateLkeRekapForm ...
type CreateLkeRekapForm struct {
	IDOPD          int64   `form:"id_opd" json:"id_opd" binding:"required"`
	Tahun          int     `form:"tahun" json:"tahun" binding:"required"`
	NilaiCapaian   float64 `form:"nilai_capaian" json:"nilai_capaian" binding:"required"`
	Kelengkapan    float64 `form:"kelengkapan" json:"kelengkapan" binding:"required"`
	PredikatAkhir  string  `form:"predikat_akhir" json:"predikat_akhir" binding:"required,max=50"`
	Predikat       string  `form:"predikat" json:"predikat" binding:"required,max=50"`
	StatusEvaluasi string  `form:"status_evaluasi" json:"status_evaluasi" binding:"required,max=50"`
	IDVerifikator  int64   `form:"id_verifikator" json:"id_verifikator" binding:"required"`
	IDKetua        int64   `form:"id_ketua" json:"id_ketua" binding:"required"`
	IDEvaluator    int64   `form:"id_evaluator" json:"id_evaluator" binding:"required"`
	IDPengendali   int64   `form:"id_pengendali" json:"id_pengendali" binding:"required"`
}

// UpdateLkeRekapForm ...
type UpdateLkeRekapForm struct {
	IDOPD          *int64   `form:"id_opd" json:"id_opd"`
	Tahun          *int     `form:"tahun" json:"tahun"`
	NilaiCapaian   *float64 `form:"nilai_capaian" json:"nilai_capaian"`
	Kelengkapan    *float64 `form:"kelengkapan" json:"kelengkapan"`
	PredikatAkhir  *string  `form:"predikat_akhir" json:"predikat_akhir" binding:"omitempty,max=50"`
	Predikat       *string  `form:"predikat" json:"predikat" binding:"omitempty,max=50"`
	StatusEvaluasi *string  `form:"status_evaluasi" json:"status_evaluasi" binding:"omitempty,max=50"`
	IDVerifikator  *int64   `form:"id_verifikator" json:"id_verifikator"`
	IDKetua        *int64   `form:"id_ketua" json:"id_ketua"`
	IDEvaluator    *int64   `form:"id_evaluator" json:"id_evaluator"`
	IDPengendali   *int64   `form:"id_pengendali" json:"id_pengendali"`
}

// IDOPD ...
func (f LkeRekapForm) IDOPD(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the ID OPD"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Tahun ...
func (f LkeRekapForm) Tahun(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the tahun"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// NilaiCapaian ...
func (f LkeRekapForm) NilaiCapaian(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the nilai capaian"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Kelengkapan ...
func (f LkeRekapForm) Kelengkapan(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the kelengkapan"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// PredikatAkhir ...
func (f LkeRekapForm) PredikatAkhir(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the predikat akhir"
		}
		return errMsg[0]
	case "max":
		return "Predikat akhir should be maximum 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Predikat ...
func (f LkeRekapForm) Predikat(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the predikat"
		}
		return errMsg[0]
	case "max":
		return "Predikat should be maximum 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// StatusEvaluasi ...
func (f LkeRekapForm) StatusEvaluasi(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the status evaluasi"
		}
		return errMsg[0]
	case "max":
		return "Status evaluasi should be maximum 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// IDVerifikator ...
func (f LkeRekapForm) IDVerifikator(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the ID verifikator"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// IDKetua ...
func (f LkeRekapForm) IDKetua(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the ID ketua"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// IDEvaluator ...
func (f LkeRekapForm) IDEvaluator(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the ID evaluator"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// IDPengendali ...
func (f LkeRekapForm) IDPengendali(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the ID pengendali"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

// Create ...
func (f LkeRekapForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "IDOPD":
				return f.IDOPD(err.Tag())
			case "Tahun":
				return f.Tahun(err.Tag())
			case "NilaiCapaian":
				return f.NilaiCapaian(err.Tag())
			case "Kelengkapan":
				return f.Kelengkapan(err.Tag())
			case "PredikatAkhir":
				return f.PredikatAkhir(err.Tag())
			case "Predikat":
				return f.Predikat(err.Tag())
			case "StatusEvaluasi":
				return f.StatusEvaluasi(err.Tag())
			case "IDVerifikator":
				return f.IDVerifikator(err.Tag())
			case "IDKetua":
				return f.IDKetua(err.Tag())
			case "IDEvaluator":
				return f.IDEvaluator(err.Tag())
			case "IDPengendali":
				return f.IDPengendali(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f LkeRekapForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "IDOPD":
				return f.IDOPD(err.Tag())
			case "Tahun":
				return f.Tahun(err.Tag())
			case "NilaiCapaian":
				return f.NilaiCapaian(err.Tag())
			case "Kelengkapan":
				return f.Kelengkapan(err.Tag())
			case "PredikatAkhir":
				return f.PredikatAkhir(err.Tag())
			case "Predikat":
				return f.Predikat(err.Tag())
			case "StatusEvaluasi":
				return f.StatusEvaluasi(err.Tag())
			case "IDVerifikator":
				return f.IDVerifikator(err.Tag())
			case "IDKetua":
				return f.IDKetua(err.Tag())
			case "IDEvaluator":
				return f.IDEvaluator(err.Tag())
			case "IDPengendali":
				return f.IDPengendali(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
