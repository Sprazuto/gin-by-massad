package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/forms"
)

// LkeEvaluasi represents the lke_evaluasi table
type LkeEvaluasi struct {
	ID             int64      `db:"id, primarykey, autoincrement" json:"id"`
	LkeRekapID     int64      `db:"lke_rekap_id" json:"lke_rekap_id"`
	UserID         int64      `db:"user_id" json:"-"`
	KodeEvaluasi   string     `db:"kode_evaluasi" json:"kode_evaluasi"`
	Jawaban        NullString `db:"jawaban" json:"jawaban"`
	Berkas         NullString `db:"berkas" json:"berkas"`
	Catatan        NullString `db:"catatan" json:"catatan"`
	KomponenBobot  float64    `db:"komponen_bobot" json:"komponen_bobot"`
	KomponenNama   string     `db:"komponen_nama" json:"komponen_nama"`
	KomponenEviden string     `db:"komponen_eviden" json:"komponen_eviden"`
	KomponenLevel  string     `db:"komponen_level" json:"komponen_level"`

	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	User      *JSONRaw `db:"user" json:"user"`
}

// LkeEvaluasiModel handles database operations
type LkeEvaluasiModel struct{}

// CreateOrUpdate checks for existing records and creates or updates accordingly
func (m LkeEvaluasiModel) CreateOrUpdate(userID int64, form forms.CreateLkeEvaluasiForm) (lkeEvaluasiID int64, err error) {
	// Check if the record exists
	var existingID int64
	err = db.GetDB().QueryRow(`
		SELECT id FROM public.lke_evaluasi
		WHERE lke_rekap_id = $1 AND kode_evaluasi = $2`,
		form.LkeRekapID, form.KodeEvaluasi).Scan(&existingID)

	if err == nil {
		// Record exists, update it
		return existingID, m.Update(userID, existingID, form)
	}

	// Record does not exist, create a new one
	return m.Create(userID, form)
}

// Create a new lke_evaluasi record
func (m LkeEvaluasiModel) Create(userID int64, form forms.CreateLkeEvaluasiForm) (lkeEvaluasiID int64, err error) {
	err = db.GetDB().QueryRow(
		`INSERT INTO public.lke_evaluasi(
			lke_rekap_id, user_id, kode_evaluasi,
			jawaban, berkas, catatan
		) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		form.LkeRekapID, userID, form.KodeEvaluasi,
		form.Jawaban, form.Berkas, form.Catatan,
	).Scan(&lkeEvaluasiID)
	return lkeEvaluasiID, err
}

// One gets a single lke_evaluasi record
func (m LkeEvaluasiModel) One(userID, id int64) (lkeEvaluasi LkeEvaluasi, err error) {
	err = db.GetDB().SelectOne(&lkeEvaluasi, `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_evaluasi l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.user_id=$1 AND l.id=$2 LIMIT 1`,
		userID, id)
	return lkeEvaluasi, err
}

// All gets all lke_evaluasi records for a user
func (m LkeEvaluasiModel) All(userID int64) (lkeEvaluasis []DataList, err error) {
	_, err = db.GetDB().Select(&lkeEvaluasis, `
		SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data,
		(SELECT row_to_json(n) FROM (
			SELECT count(l.id) AS total
			FROM public.lke_evaluasi AS l
			WHERE l.user_id=$1 LIMIT 1
		) n ) AS meta
		FROM (
			SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
			FROM public.lke_evaluasi l
			LEFT JOIN public.user u ON l.user_id = u.id
			WHERE l.user_id=$1
			ORDER BY l.id DESC
		) d`,
		userID)
	return lkeEvaluasis, err
}

// Update an existing lke_evaluasi record
func (m LkeEvaluasiModel) Update(userID int64, id int64, form forms.CreateLkeEvaluasiForm) (err error) {
	// Validate required fields
	if form.LkeRekapID == 0 {
		return errors.New("lke_rekap_id is required")
	}
	if form.KodeEvaluasi == "" {
		return errors.New("kode_evaluasi is required")
	}

	// Build dynamic SQL query based on field presence
	query := "UPDATE public.lke_evaluasi SET"
	var args []interface{}
	argCount := 1

	// Always include required fields
	query += fmt.Sprintf(" lke_rekap_id=$%d, kode_evaluasi=$%d", argCount, argCount+1)
	args = append(args, form.LkeRekapID, form.KodeEvaluasi)
	argCount += 2

	// Add fields only if they are present in the form
	if form.Jawaban != nil {
		query += fmt.Sprintf(", jawaban=$%d", argCount)
		args = append(args, *form.Jawaban)
		argCount++
	}
	if form.Berkas != nil {
		query += fmt.Sprintf(", berkas=$%d", argCount)
		args = append(args, *form.Berkas)
		argCount++
	}
	// Always update Catatan if present in request
	if form.Catatan != nil {
		query += fmt.Sprintf(", catatan=$%d", argCount)
		args = append(args, *form.Catatan)
		argCount++
	} else {
		// If Catatan is not in request, keep existing value
		query += ", catatan=catatan"
	}

	// Add WHERE clause
	query += " WHERE id=$" + strconv.Itoa(argCount)
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

// Delete an lke_evaluasi record
func (m LkeEvaluasiModel) Delete(userID, id int64) (err error) {
	operation, err := db.GetDB().Exec("DELETE FROM public.lke_evaluasi WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}

// CalculateKelengkapan calculates the completeness percentage
func (m LkeEvaluasiModel) CalculateKelengkapan(lkeRekapID int64) (float64, error) {
	var kelengkapan float64
	var evalCount int
	var komponenCount int

	// Get count of lke_evaluasi for this rekap
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM lke_evaluasi WHERE lke_rekap_id = $1", lkeRekapID).Scan(&evalCount)
	if err != nil {
		return 0, err
	}

	// Get count of lke_komponen with bobot > 0
	err = db.GetDB().QueryRow("SELECT COUNT(*) FROM lke_komponen WHERE bobot > 0").Scan(&komponenCount)
	if err != nil {
		return 0, err
	}

	// Calculate kelengkapan
	if komponenCount > 0 {
		kelengkapan = (float64(evalCount) / float64(komponenCount)) * 100
	}

	return kelengkapan, nil
}

// CalculateNilaiCapaian calculates the achievement score
func (m LkeEvaluasiModel) CalculateNilaiCapaian(lkeRekapID int64) (float64, error) {
	var nilaiCapaian float64

	rows, err := db.GetDB().Query(`
		SELECT e.jawaban, k.bobot
		FROM lke_evaluasi e
		JOIN lke_komponen k ON e.kode_evaluasi = k.kode_evaluasi
		WHERE e.lke_rekap_id = $1`, lkeRekapID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var jawaban string
		var bobot float64
		if err := rows.Scan(&jawaban, &bobot); err != nil {
			continue
		}

		switch jawaban {
		case "Ya", "Sudah":
			nilaiCapaian += bobot
		case "Sebagian":
			nilaiCapaian += bobot / 2
		case "Belum", "Tidak":
			nilaiCapaian += 0
		}
	}

	return nilaiCapaian, nil
}

// UpdateRekapValues updates kelengkapan and nilai_capaian in lke_rekap
func (m LkeEvaluasiModel) UpdateRekapValues(lkeRekapID int64) error {
	kelengkapan, err := m.CalculateKelengkapan(lkeRekapID)
	if err != nil {
		return fmt.Errorf("failed to calculate kelengkapan: %w", err)
	}

	nilaiCapaian, err := m.CalculateNilaiCapaian(lkeRekapID)
	if err != nil {
		return fmt.Errorf("failed to calculate nilai_capaian: %w", err)
	}

	_, err = db.GetDB().Exec(`
		UPDATE lke_rekap
		SET kelengkapan = $1, nilai_capaian = $2
		WHERE id = $3`,
		kelengkapan, nilaiCapaian, lkeRekapID)
	if err != nil {
		return fmt.Errorf("failed to update lke_rekap: %w", err)
	}
	return nil
}
