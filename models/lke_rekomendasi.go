package models

import (
	"errors"

	"lke-app/db"
	"lke-app/forms"
)

// LkeRekomendasi represents the lke_rekomendasi table
type LkeRekomendasi struct {
	ID       int64      `db:"id, primarykey, autoincrement" json:"id"`
	ParentID int64      `db:"parent_id" json:"parent_id"`
	Ta1a     NullString `db:"ta1a" json:"ta1a"`
	Ta1b     NullString `db:"ta1b" json:"ta1b"`
	Ta1c     NullString `db:"ta1c" json:"ta1c"`
	Ta2a     NullString `db:"ta2a" json:"ta2a"`
	Ta2b     NullString `db:"ta2b" json:"ta2b"`
	Ta2c     NullString `db:"ta2c" json:"ta2c"`
	Tb1      NullString `db:"tb1" json:"tb1"`
	Tb2      NullString `db:"tb2" json:"tb2"`
	Tb3      NullString `db:"tb3" json:"tb3"`
	Tc1      NullString `db:"tc1" json:"tc1"`
	Tc2      NullString `db:"tc2" json:"tc2"`
	Tc3      NullString `db:"tc3" json:"tc3"`
	Td1      NullString `db:"td1" json:"td1"`
	Td2      NullString `db:"td2" json:"td2"`
	Td3      NullString `db:"td3" json:"td3"`
}

// LkeRekomendasiModel ...
type LkeRekomendasiModel struct{}

// GetByParentID ...
func (m LkeRekomendasiModel) GetByParentID(parentID int64) (rekomendasi LkeRekomendasi, err error) {
	err = db.GetDB().SelectOne(&rekomendasi, "SELECT id, parent_id, ta1a, ta1b, ta1c, ta2a, ta2b, ta2c, tb1, tb2, tb3, tc1, tc2, tc3, td1, td2, td3 FROM public.lke_rekomendasi WHERE parent_id=$1 LIMIT 1", parentID)
	return rekomendasi, err
}

// Create ...
func (m LkeRekomendasiModel) Create(form forms.CreateLkeRekomendasiForm) (rekomendasiID int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.lke_rekomendasi(parent_id, ta1a, ta1b, ta1c, ta2a, ta2b, ta2c, tb1, tb2, tb3, tc1, tc2, tc3, td1, td2, td3) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id", form.ParentID, form.Ta1a, form.Ta1b, form.Ta1c, form.Ta2a, form.Ta2b, form.Ta2c, form.Tb1, form.Tb2, form.Tb3, form.Tc1, form.Tc2, form.Tc3, form.Td1, form.Td2, form.Td3).Scan(&rekomendasiID)
	return rekomendasiID, err
}

// Upsert ...
func (m LkeRekomendasiModel) Upsert(form forms.CreateLkeRekomendasiForm) (rekomendasiID int64, isUpdate bool, err error) {
	// Check if record exists with the same parent_id
	existing, err := m.GetByParentID(form.ParentID)
	if err != nil {
		// If no record found, create new one
		if err.Error() == "sql: no rows in result set" {
			rekomendasiID, err = m.Create(form)
			return rekomendasiID, false, err
		}
		// Other errors
		return 0, false, err
	}

	// Record exists, update it
	err = m.Update(existing.ID, form)
	if err != nil {
		return 0, false, err
	}

	return existing.ID, true, nil
}

// One ...
func (m LkeRekomendasiModel) One(id int64) (rekomendasi LkeRekomendasi, err error) {
	err = db.GetDB().SelectOne(&rekomendasi, "SELECT id, parent_id, ta1a, ta1b, ta1c, ta2a, ta2b, ta2c, tb1, tb2, tb3, tc1, tc2, tc3, td1, td2, td3 FROM public.lke_rekomendasi WHERE id=$1 LIMIT 1", id)
	return rekomendasi, err
}

// All ...
func (m LkeRekomendasiModel) All() (rekomendasi []LkeRekomendasi, err error) {
	_, err = db.GetDB().Select(&rekomendasi, "SELECT id, parent_id, ta1a, ta1b, ta1c, ta2a, ta2b, ta2c, tb1, tb2, tb3, tc1, tc2, tc3, td1, td2, td3 FROM public.lke_rekomendasi ORDER BY id DESC")
	return rekomendasi, err
}

// Update ...
func (m LkeRekomendasiModel) Update(id int64, form forms.CreateLkeRekomendasiForm) (err error) {
	operation, err := db.GetDB().Exec("UPDATE public.lke_rekomendasi SET parent_id=$2, ta1a=$3, ta1b=$4, ta1c=$5, ta2a=$6, ta2b=$7, ta2c=$8, tb1=$9, tb2=$10, tb3=$11, tc1=$12, tc2=$13, tc3=$14, td1=$15, td2=$16, td3=$17 WHERE id=$1", id, form.ParentID, form.Ta1a, form.Ta1b, form.Ta1c, form.Ta2a, form.Ta2b, form.Ta2c, form.Tb1, form.Tb2, form.Tb3, form.Tc1, form.Tc2, form.Tc3, form.Td1, form.Td2, form.Td3)
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
func (m LkeRekomendasiModel) Delete(id int64) (err error) {
	operation, err := db.GetDB().Exec("DELETE FROM public.lke_rekomendasi WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
