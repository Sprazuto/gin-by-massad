package models

import (
	"errors"

	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/forms"
)

// LkeKomponen represents the lke_komponen table
type LkeKomponen struct {
	ID           int64      `db:"id, primarykey, autoincrement" json:"id"`
	KodeEvaluasi string     `db:"kode_evaluasi" json:"kode_evaluasi"`
	Bobot        float64    `db:"bobot" json:"bobot"`
	Komponen     NullString `db:"komponen" json:"komponen"`
	Eviden       NullString `db:"eviden" json:"eviden"`
	Level        NullString `db:"level" json:"level"`
	CreatedAt    int64      `db:"created_at" json:"created_at"`
	UpdatedAt    int64      `db:"updated_at" json:"updated_at"`
}

// LkeKomponenModel ...
type LkeKomponenModel struct{}

// Create ...
func (m LkeKomponenModel) Create(form forms.CreateLkeKomponenForm) (komponenID int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.lke_komponen(kode_evaluasi, bobot, komponen, eviden, level) VALUES($1, $2, $3, $4, $5) RETURNING id", form.KodeEvaluasi, form.Bobot, form.Komponen, form.Eviden, form.Level).Scan(&komponenID)
	return komponenID, err
}

// One ...
func (m LkeKomponenModel) One(id int64) (komponen LkeKomponen, err error) {
	err = db.GetDB().SelectOne(&komponen, "SELECT id, kode_evaluasi, bobot, komponen, eviden, level, created_at, updated_at FROM public.lke_komponen WHERE id=$1 LIMIT 1", id)
	return komponen, err
}

// All ...
func (m LkeKomponenModel) All() (komponen []LkeKomponen, err error) {
	_, err = db.GetDB().Select(&komponen, "SELECT id, kode_evaluasi, bobot, komponen, eviden, level, created_at, updated_at FROM public.lke_komponen ORDER BY id DESC")
	return komponen, err
}

// Update ...
func (m LkeKomponenModel) Update(id int64, form forms.CreateLkeKomponenForm) (err error) {
	operation, err := db.GetDB().Exec("UPDATE public.lke_komponen SET kode_evaluasi=$2, bobot=$3, komponen=$4, eviden=$5, level=$6 WHERE id=$1", id, form.KodeEvaluasi, form.Bobot, form.Komponen, form.Eviden, form.Level)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

// Delete ...
func (m LkeKomponenModel) Delete(id int64) (err error) {
	operation, err := db.GetDB().Exec("DELETE FROM public.lke_komponen WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
