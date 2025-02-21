package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/forms"
)

// LkeRekap represents the lke_rekap table
type LkeRekap struct {
	ID             int64         `db:"id, primarykey, autoincrement" json:"id"`
	UserID         int64         `db:"user_id" json:"-"`
	IDOPD          int64         `db:"id_opd" json:"id_opd"`
	Tahun          int           `db:"tahun" json:"tahun"`
	NilaiCapaian   float64       `db:"nilai_capaian" json:"nilai_capaian"`
	Kelengkapan    float64       `db:"kelengkapan" json:"kelengkapan"`
	PredikatAkhir  string        `db:"predikat_akhir" json:"predikat_akhir"`
	Predikat       string        `db:"predikat" json:"predikat"`
	StatusEvaluasi string        `db:"status_evaluasi" json:"status_evaluasi"`
	IDVerifikator  NullInt64     `db:"id_verifikator" json:"id_verifikator"`
	IDKetua        NullInt64     `db:"id_ketua" json:"id_ketua"`
	IDEvaluator    NullInt64     `db:"id_evaluator" json:"id_evaluator"`
	IDPengendali   NullInt64     `db:"id_pengendali" json:"id_pengendali"`
	UpdatedAt      int64         `db:"updated_at" json:"updated_at"`
	CreatedAt      int64         `db:"created_at" json:"created_at"`
	User           *JSONRaw      `db:"user" json:"user"`
	Evaluasi       []LkeEvaluasi `json:"evaluasi"`
}

// LkeRekapModel handles database operations
type LkeRekapModel struct{}

// Create a new lke_rekap record
func (m LkeRekapModel) Create(userID int64, form forms.CreateLkeRekapForm) (lkeRekapID int64, err error) {
	err = db.GetDB().QueryRow(
		`INSERT INTO public.lke_rekap(
			user_id, id_opd, tahun, nilai_capaian, kelengkapan,
			predikat_akhir, predikat, status_evaluasi,
			id_verifikator, id_ketua, id_evaluator, id_pengendali
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`,
		userID, form.IDOPD, form.Tahun, form.NilaiCapaian, form.Kelengkapan,
		form.PredikatAkhir, form.Predikat, form.StatusEvaluasi,
		form.IDVerifikator, form.IDKetua, form.IDEvaluator, form.IDPengendali,
	).Scan(&lkeRekapID)
	return lkeRekapID, err
}

// One gets a single lke_rekap record
func (m LkeRekapModel) One(userID, id int64) (lkeRekap LkeRekap, err error) {
	err = db.GetDB().SelectOne(&lkeRekap, `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_rekap l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.user_id=$1 AND l.id=$2 LIMIT 1`,
		userID, id)
	return lkeRekap, err
}

// All gets all lke_rekap records for a user
func (m LkeRekapModel) All(userID int64) (lkeRekaps []DataList, err error) {
	_, err = db.GetDB().Select(&lkeRekaps, `
		SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data,
		(SELECT row_to_json(n) FROM (
			SELECT count(l.id) AS total
			FROM public.lke_rekap AS l
			WHERE l.user_id=$1 LIMIT 1
		) n ) AS meta
		FROM (
			SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
			FROM public.lke_rekap l
			LEFT JOIN public.user u ON l.user_id = u.id
			WHERE l.user_id=$1
			ORDER by l.id DESC
		) d`,
		userID)
	return lkeRekaps, err
}

// Update an existing lke_rekap record
func (m LkeRekapModel) Update(userID int64, id int64, form forms.UpdateLkeRekapForm) (err error) {
	// Build dynamic SQL query based on non-nil fields
	query := "UPDATE public.lke_rekap SET"
	var args []interface{}
	argCount := 1

	if form.IDOPD != nil {
		query += fmt.Sprintf(" id_opd=$%d,", argCount)
		args = append(args, *form.IDOPD)
		argCount++
	}
	if form.Tahun != nil {
		query += fmt.Sprintf(" tahun=$%d,", argCount)
		args = append(args, *form.Tahun)
		argCount++
	}
	if form.NilaiCapaian != nil {
		query += fmt.Sprintf(" nilai_capaian=$%d,", argCount)
		args = append(args, *form.NilaiCapaian)
		argCount++
	}
	if form.Kelengkapan != nil {
		query += fmt.Sprintf(" kelengkapan=$%d,", argCount)
		args = append(args, *form.Kelengkapan)
		argCount++
	}
	if form.PredikatAkhir != nil {
		query += fmt.Sprintf(" predikat_akhir=$%d,", argCount)
		args = append(args, *form.PredikatAkhir)
		argCount++
	}
	if form.Predikat != nil {
		query += fmt.Sprintf(" predikat=$%d,", argCount)
		args = append(args, *form.Predikat)
		argCount++
	}
	if form.StatusEvaluasi != nil {
		query += fmt.Sprintf(" status_evaluasi=$%d,", argCount)
		args = append(args, *form.StatusEvaluasi)
		argCount++
	}
	if form.IDVerifikator != nil {
		query += fmt.Sprintf(" id_verifikator=$%d,", argCount)
		args = append(args, *form.IDVerifikator)
		argCount++
	}
	if form.IDKetua != nil {
		query += fmt.Sprintf(" id_ketua=$%d,", argCount)
		args = append(args, *form.IDKetua)
		argCount++
	}
	if form.IDEvaluator != nil {
		query += fmt.Sprintf(" id_evaluator=$%d,", argCount)
		args = append(args, *form.IDEvaluator)
		argCount++
	}
	if form.IDPengendali != nil {
		query += fmt.Sprintf(" id_pengendali=$%d,", argCount)
		args = append(args, *form.IDPengendali)
		argCount++
	}

	// Remove trailing comma and add WHERE clause
	query = strings.TrimSuffix(query, ",") + " WHERE id=$" + strconv.Itoa(argCount)
	args = append(args, id)

	operation, err := db.GetDB().Exec(query, args...)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

// Delete an lke_rekap record
func (m LkeRekapModel) Delete(userID, id int64) (err error) {
	operation, err := db.GetDB().Exec("DELETE FROM public.lke_rekap WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}

// OneWithEvaluasi gets a lke_rekap record by id_opd and tahun with its evaluasi children
func (m LkeRekapModel) OneWithEvaluasi(userID int64, idOPD int64, tahun int) (lkeRekap LkeRekap, err error) {
	// Get the lke_rekap record
	query := `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_rekap l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.id_opd=$1 AND l.tahun=$2 LIMIT 1`

	err = db.GetDB().SelectOne(&lkeRekap, query, idOPD, tahun)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Insert default data if record doesn't exist
			defaultForm := forms.CreateLkeRekapForm{
				IDOPD:          idOPD,
				Tahun:          tahun,
				NilaiCapaian:   0,
				Kelengkapan:    0,
				PredikatAkhir:  "-",
				Predikat:       "-",
				StatusEvaluasi: "Belum Dievaluasi",
			}
			_, err = m.Create(userID, defaultForm)
			if err != nil {
				return lkeRekap, err
			}

			// Retry fetching the newly created record
			err = db.GetDB().SelectOne(&lkeRekap, query, idOPD, tahun)
			if err != nil {
				return lkeRekap, err
			}
			return lkeRekap, nil
		}
		return lkeRekap, err
	}

	// Get related lke_evaluasi records with lke_komponen join (this can fail silently)
	_, _ = db.GetDB().Select(&lkeRekap.Evaluasi, `
		SELECT e.*,
		       k.bobot as komponen_bobot,
		       k.komponen as komponen_nama,
		       k.eviden as komponen_eviden,
		       k.level as komponen_level
		FROM public.lke_evaluasi e
		LEFT JOIN public.lke_komponen k ON e.kode_evaluasi = k.kode_evaluasi
		WHERE e.lke_rekap_id=$1`,
		lkeRekap.ID)

	return lkeRekap, nil
}
