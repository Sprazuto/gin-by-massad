package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"lke-app/db"
	"lke-app/forms"
)

// Status constants for LKE Rekap workflow
const (
	StatusDraft            = "draft"
	StatusSekdisReview     = "sekdis_review"
	StatusEvaluatorReview  = "evaluator_review"
	StatusKetuaReview      = "ketua_review"
	StatusPengendaliReview = "pengendali_review"
	StatusIrbanReview      = "irban_review"
	StatusFinal            = "final"
)

// validTransitions defines allowed status transitions
// Format: current_status -> []allowed_next_statuses
var validTransitions = map[string][]string{
	StatusDraft:            {StatusSekdisReview},
	StatusSekdisReview:     {StatusEvaluatorReview, StatusDraft},
	StatusEvaluatorReview:  {StatusKetuaReview, StatusDraft},
	StatusKetuaReview:      {StatusPengendaliReview, StatusEvaluatorReview},
	StatusPengendaliReview: {StatusIrbanReview, StatusEvaluatorReview},
	StatusIrbanReview:      {StatusFinal, StatusEvaluatorReview},
}

// LkeRekap represents the lke_rekap table
type LkeRekap struct {
	ID                int64           `db:"id, primarykey, autoincrement" json:"id"`
	UserID            int64           `db:"user_id" json:"-"`
	IDOPD             int64           `db:"id_opd" json:"id_opd"`
	Tahun             int             `db:"tahun" json:"tahun"`
	Kelengkapan       NullFloat64     `db:"kelengkapan" json:"kelengkapan"`
	NilaiCapaian      NullFloat64     `db:"nilai_capaian" json:"nilai_capaian"`
	PredikatAkhir     NullString      `db:"predikat_akhir" json:"predikat_akhir"`
	Predikat          NullString      `db:"predikat" json:"predikat"`
	KelengkapanM      NullFloat64     `db:"kelengkapan_m" json:"kelengkapan_m"`
	NilaiCapaianM     NullFloat64     `db:"nilai_capaian_m" json:"nilai_capaian_m"`
	PredikatAkhirM    NullString      `db:"predikat_akhir_m" json:"predikat_akhir_m"`
	PredikatM         NullString      `db:"predikat_m" json:"predikat_m"`
	StatusEvaluasi    string          `db:"status_evaluasi" json:"status_evaluasi"`
	StatusLabel       string          `db:"status_label" json:"status_label"`
	PreviousStatus    NullString      `db:"previous_status" json:"previous_status"`
	StatusDescription NullString      `db:"status_description" json:"status_description"`
	IDSekdis          NullInt64       `db:"id_sekdis" json:"id_sekdis"`
	IDVerifikator     NullInt64       `db:"id_verifikator" json:"id_verifikator"`
	IDKetua           NullInt64       `db:"id_ketua" json:"id_ketua"`
	IDEvaluator       NullInt64       `db:"id_evaluator" json:"id_evaluator"`
	IDPengendali      NullInt64       `db:"id_pengendali" json:"id_pengendali"`
	IDIrban           NullInt64       `db:"id_irban" json:"id_irban"`
	UpdatedAt         int64           `db:"updated_at" json:"updated_at"`
	CreatedAt         int64           `db:"created_at" json:"created_at"`
	User              *JSONRaw        `db:"user" json:"user"`
	Evaluasi          []LkeEvaluasi   `json:"evaluasi"`
	Rekomendasi       *LkeRekomendasi `json:"rekomendasi"`
}

// StatusLabelMap maps internal status to user-readable Indonesian labels
var StatusLabelMap = map[string]string{
	StatusDraft:            "Belum Dievaluasi",
	StatusSekdisReview:     "Review Sekdis",
	StatusEvaluatorReview:  "Review Evaluator",
	StatusKetuaReview:      "Review Ketua",
	StatusPengendaliReview: "Review Pengendali",
	StatusIrbanReview:      "Review IRBAN",
	StatusFinal:            "Sudah Dievaluasi",
}

// GetStatusLabel returns the localized label for a status
func GetStatusLabel(status string) string {
	if label, ok := StatusLabelMap[status]; ok {
		return label
	}
	return status
}

// LkeRekapModel handles database operations
type LkeRekapModel struct{}

// IsValidTransition checks if status transition is allowed
func (m LkeRekapModel) IsValidTransition(currentStatus, newStatus string) bool {
	allowed, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}
	for _, s := range allowed {
		if s == newStatus {
			return true
		}
	}
	return false
}

// StatusDescriptionMap provides human-readable descriptions based on previous status
// Used when record transitions TO a new status
var StatusDescriptionMap = map[string]string{
	StatusDraft:            "Belum dievaluasi",
	StatusSekdisReview:     "Draft dikembalikan oleh Sekretaris untuk perbaikan",
	StatusEvaluatorReview:  "Draft dikembalikan oleh Evaluator untuk perbaikan",
	StatusKetuaReview:      "Evaluasi ulang diperlukan atas permintaan dari Ketua Tim",
	StatusPengendaliReview: "Evaluasi ulang diperlukan atas permintaan dari Pengendali Teknis",
	StatusIrbanReview:      "Evaluasi ulang diperlukan atas permintaan dari Inspektur Pembantu",
}

// StatusForwardDescriptionMap provides human-readable descriptions based on new status
// Used for forward/approval transitions
var StatusForwardDescriptionMap = map[string]string{
	StatusDraft:            "Belum dievaluasi",
	StatusSekdisReview:     "Diteruskan ke Sekretaris untuk review",
	StatusEvaluatorReview:  "Diteruskan ke Evaluator oleh Sekretaris",
	StatusKetuaReview:      "Diteruskan ke Ketua Tim oleh Evaluator",
	StatusPengendaliReview: "Diteruskan ke Pengendali Teknis oleh Ketua Tim",
	StatusIrbanReview:      "Diteruskan ke Inspektur Pembantu oleh Pengendali Teknis",
	StatusFinal:            "Sudah dievaluasi",
}

// ResetDescription is used when a record is manually reset to draft
const ResetDescription = "Direset manual oleh admin"

// GetStatusDescription returns human-readable description based on previous status
func GetStatusDescription(previousStatus string) string {
	if desc, ok := StatusDescriptionMap[previousStatus]; ok {
		return desc
	}
	return ""
}

// GetForwardDescription returns human-readable description based on new status
func GetForwardDescription(newStatus string) string {
	if desc, ok := StatusForwardDescriptionMap[newStatus]; ok {
		return desc
	}
	return ""
}

// UpdateStatus transitions the status with validation and updates status_label
// It also tracks previous_status and status_description for all transitions
func (m LkeRekapModel) UpdateStatus(id int64, newStatus string) error {
	// Get current record
	var currentStatus string
	err := db.GetDB().QueryRow(
		"SELECT status_evaluasi FROM public.lke_rekap WHERE id=$1", id,
	).Scan(&currentStatus)
	if err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	// Validate transition
	if !m.IsValidTransition(currentStatus, newStatus) {
		return fmt.Errorf("invalid status transition from %s to %s", currentStatus, newStatus)
	}

	// Get status label
	newStatusLabel := GetStatusLabel(newStatus)

	// Track previous status
	previousStatus := currentStatus

	// Determine status_description based on transition direction
	// Backward transitions (rejecting to earlier stage) use GetStatusDescription
	// Forward transitions (approving to next stage) use GetForwardDescription
	var statusDescription string
	if isBackwardTransition(currentStatus, newStatus) {
		statusDescription = GetStatusDescription(currentStatus)
	} else {
		statusDescription = GetForwardDescription(newStatus)
	}

	// Update status, label, previous_status, and description
	_, err = db.GetDB().Exec(
		"UPDATE public.lke_rekap SET status_evaluasi=$1, status_label=$2, previous_status=$3, status_description=$4, updated_at=$5 WHERE id=$6",
		newStatus, newStatusLabel, previousStatus, statusDescription, time.Now().Unix(), id,
	)
	return err
}

// isBackwardTransition checks if a transition is going backwards in the workflow
func isBackwardTransition(currentStatus, newStatus string) bool {
	// Define status order (higher index = later stage)
	statusOrder := []string{
		StatusDraft,
		StatusSekdisReview,
		StatusEvaluatorReview,
		StatusKetuaReview,
		StatusPengendaliReview,
		StatusIrbanReview,
		StatusFinal,
	}

	currentIdx := -1
	newIdx := -1
	for i, s := range statusOrder {
		if s == currentStatus {
			currentIdx = i
		}
		if s == newStatus {
			newIdx = i
		}
	}

	// Backward if new status comes earlier in the workflow
	return newIdx < currentIdx
}

// UpdateStatusWithDescription transitions the status with a custom description
// This is useful when you want to set a specific description for the transition
func (m LkeRekapModel) UpdateStatusWithDescription(id int64, newStatus, description string) error {
	// Get current record
	var currentStatus string
	err := db.GetDB().QueryRow(
		"SELECT status_evaluasi FROM public.lke_rekap WHERE id=$1", id,
	).Scan(&currentStatus)
	if err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	// Validate transition
	if !m.IsValidTransition(currentStatus, newStatus) {
		return fmt.Errorf("invalid status transition from %s to %s", currentStatus, newStatus)
	}

	// Get status label
	newStatusLabel := GetStatusLabel(newStatus)

	// Track previous status
	previousStatus := currentStatus

	// Use provided description or fallback to forward description
	statusDescription := description
	if statusDescription == "" {
		statusDescription = GetForwardDescription(newStatus)
	}

	// Update status, label, previous_status, and description
	_, err = db.GetDB().Exec(
		"UPDATE public.lke_rekap SET status_evaluasi=$1, status_label=$2, previous_status=$3, status_description=$4, updated_at=$5 WHERE id=$6",
		newStatus, newStatusLabel, previousStatus, statusDescription, time.Now().Unix(), id,
	)
	return err
}

// ResetToDraft resets any record back to draft (admin function)
// Allows resetting from any status to draft with reset description
func (m LkeRekapModel) ResetToDraft(id int64) error {
	// Get current record
	var currentStatus string
	err := db.GetDB().QueryRow(
		"SELECT status_evaluasi FROM public.lke_rekap WHERE id=$1", id,
	).Scan(&currentStatus)
	if err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	// Cannot reset if already in draft
	if currentStatus == StatusDraft {
		return fmt.Errorf("record is already in draft status")
	}

	// Update to draft with reset description
	_, err = db.GetDB().Exec(
		"UPDATE public.lke_rekap SET status_evaluasi=$1, status_label=$2, previous_status=$3, status_description=$4, updated_at=$5 WHERE id=$6",
		StatusDraft, GetStatusLabel(StatusDraft), currentStatus, ResetDescription, time.Now().Unix(), id,
	)
	return err
}

// GetByStatus retrieves records by status
func (m LkeRekapModel) GetByStatus(status string) ([]LkeRekap, error) {
	var results []LkeRekap
	_, err := db.GetDB().Select(&results, `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_rekap l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.status_evaluasi=$1
		ORDER BY l.id DESC`, status)
	return results, err
}

// GetOne retrieves a single lke_rekap record by ID (without user filter)
func (m LkeRekapModel) GetOne(id int64) (lkeRekap LkeRekap, err error) {
	err = db.GetDB().SelectOne(&lkeRekap, `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_rekap l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.id=$1 LIMIT 1`,
		id)
	return lkeRekap, err
}

// Create a new lke_rekap record
func (m LkeRekapModel) Create(userID int64, form forms.CreateLkeRekapForm) (lkeRekapID int64, err error) {
	// Get status label and description from status_evaluasi
	statusLabel := GetStatusLabel(form.StatusEvaluasi)
	statusDescription := GetForwardDescription(form.StatusEvaluasi)

	err = db.GetDB().QueryRow(
		`INSERT INTO public.lke_rekap(
			user_id, id_opd, tahun, nilai_capaian, kelengkapan,
			predikat_akhir, predikat, kelengkapan_m, nilai_capaian_m, predikat_akhir_m, predikat_m,
			status_evaluasi, status_label, status_description,
			id_sekdis, id_verifikator, id_ketua, id_evaluator, id_pengendali, id_irban
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id`,
		userID, form.IDOPD, form.Tahun, form.NilaiCapaian, form.Kelengkapan,
		form.PredikatAkhir, form.Predikat, form.KelengkapanM, form.NilaiCapaianM, form.PredikatAkhirM, form.PredikatM,
		form.StatusEvaluasi, statusLabel, statusDescription,
		form.IDSekdis, form.IDVerifikator, form.IDKetua, form.IDEvaluator, form.IDPengendali, form.IDIrban,
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
func (m LkeRekapModel) All(userID int64, tahun int) (lkeRekaps []DataList, err error) {
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
			WHERE l.user_id=$1 AND l.tahun=$2
			ORDER by l.id DESC
		) d`,
		userID, tahun)
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
	if form.KelengkapanM != nil {
		query += fmt.Sprintf(" kelengkapan_m=$%d,", argCount)
		args = append(args, *form.KelengkapanM)
		argCount++
	}
	if form.NilaiCapaianM != nil {
		query += fmt.Sprintf(" nilai_capaian_m=$%d,", argCount)
		args = append(args, *form.NilaiCapaianM)
		argCount++
	}
	if form.PredikatAkhirM != nil {
		query += fmt.Sprintf(" predikat_akhir_m=$%d,", argCount)
		args = append(args, *form.PredikatAkhirM)
		argCount++
	}
	if form.PredikatM != nil {
		query += fmt.Sprintf(" predikat_m=$%d,", argCount)
		args = append(args, *form.PredikatM)
		argCount++
	}
	if form.StatusEvaluasi != nil {
		query += fmt.Sprintf(" status_evaluasi=$%d,", argCount)
		args = append(args, *form.StatusEvaluasi)
		argCount++
	}
	if form.IDSekdis != nil {
		query += fmt.Sprintf(" id_sekdis=$%d,", argCount)
		args = append(args, *form.IDSekdis)
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
	if form.IDIrban != nil {
		query += fmt.Sprintf(" id_irban=$%d,", argCount)
		args = append(args, *form.IDIrban)
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

// IDSekdisGroupedResponse represents the response structure for id_sekdis grouped by status_evaluasi
type IDSekdisGroupedResponse struct {
	ByStatus    map[string][]int64 `json:"by_status"`
	AllIDSekdis []int64            `json:"all_id_sekdis"`
}

// GetIDSekdisByStatusEvaluasi gets all id_sekdis records grouped by status_evaluasi
func (m LkeRekapModel) GetIDSekdisByStatusEvaluasi(tahun int) (IDSekdisGroupedResponse, error) {
	response := IDSekdisGroupedResponse{
		ByStatus:    make(map[string][]int64),
		AllIDSekdis: []int64{},
	}

	// Get all records with non-null id_sekdis for the given tahun
	type Record struct {
		StatusEvaluasi string `db:"status_evaluasi"`
		IDSekdis       int64  `db:"id_sekdis"`
	}

	var records []Record
	_, err := db.GetDB().Select(&records, `
		SELECT status_evaluasi, id_sekdis
		FROM public.lke_rekap
		WHERE tahun = $1 AND id_sekdis IS NOT NULL
		ORDER BY id_sekdis`,
		tahun)
	if err != nil {
		return response, err
	}

	// Use map to track unique id_sekdis values
	uniqueIDs := make(map[int64]bool)

	for _, record := range records {
		// Add to by_status map
		response.ByStatus[record.StatusEvaluasi] = append(response.ByStatus[record.StatusEvaluasi], record.IDSekdis)
		// Track unique id_sekdis
		uniqueIDs[record.IDSekdis] = true
	}

	// Convert uniqueIDs map to slice
	for id := range uniqueIDs {
		response.AllIDSekdis = append(response.AllIDSekdis, id)
	}

	return response, nil
}

// OneWithEvaluasi gets a lke_rekap record by id_opd and tahun with its evaluasi children
func (m LkeRekapModel) OneWithEvaluasi(userID int64, idOPD int64, tahun int) (lkeRekap LkeRekap, err error) {
	// Get the lke_rekap record
	query := `
		SELECT l.*, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user
		FROM public.lke_rekap l
		LEFT JOIN public.user u ON l.user_id = u.id
		WHERE l.user_id=$1 AND l.id_opd=$2 AND l.tahun=$3 LIMIT 1`

	err = db.GetDB().SelectOne(&lkeRekap, query, userID, idOPD, tahun)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Insert default data if record doesn't exist
			defaultForm := forms.CreateLkeRekapForm{
				IDOPD:          idOPD,
				Tahun:          tahun,
				Kelengkapan:    0,
				NilaiCapaian:   0,
				PredikatAkhir:  "-",
				Predikat:       "-",
				NilaiCapaianM:  0,
				KelengkapanM:   0,
				PredikatAkhirM: "-",
				PredikatM:      "-",
				StatusEvaluasi: StatusDraft,
			}
			_, err = m.Create(userID, defaultForm)
			if err != nil {
				return lkeRekap, err
			}

			// Retry fetching the newly created record
			err = db.GetDB().SelectOne(&lkeRekap, query, userID, idOPD, tahun)
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
		       k.level as komponen_level,
		       CASE WHEN e.evaluasi IS NOT NULL AND e.evaluasi IN ('Ya', 'Sudah') THEN k.bobot WHEN e.evaluasi = 'Sebagian' THEN k.bobot / 2 ELSE 0 END AS capaian
		FROM public.lke_evaluasi e
		LEFT JOIN public.lke_komponen k ON e.kode_evaluasi = k.kode_evaluasi
		WHERE e.lke_rekap_id=$1`,
		lkeRekap.ID)

	// Get related lke_rekomendasi record (1:1 relationship)
	lkeRekap.Rekomendasi = &LkeRekomendasi{}
	_ = db.GetDB().SelectOne(lkeRekap.Rekomendasi, `
		SELECT r.*
		FROM public.lke_rekomendasi r
		WHERE r.parent_id=$1 LIMIT 1`,
		lkeRekap.ID)

	return lkeRekap, nil
}
