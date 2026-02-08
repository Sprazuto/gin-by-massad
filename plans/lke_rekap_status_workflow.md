# LKE Rekap Status Workflow Implementation Plan

## Overview

Transform `status_evaluasi` from a simple string field into a proper workflow status model with state machine validation and dedicated API endpoints.

## Status Constants

### Status Values

| Status            | Value               | Description                     |
| ----------------- | ------------------- | ------------------------------- |
| Draft             | `draft`             | Initial state, creator can edit |
| Sekdis Review     | `sekdis_review`     | Waiting for SEKDIS approval     |
| Evaluator Review  | `evaluator_review`  | With Evaluator for review       |
| Ketua Review      | `ketua_review`      | Waiting for KETUA approval      |
| Pengendali Review | `pengendali_review` | With Pengendali for review      |
| Irban Review      | `irban_review`      | With IRBAN for final review     |
| Final             | `final`             | Completed and finalized         |

### Status Transition Rules

| Current Status      | Next Status (Approve) | Next Status (Reject) |
| ------------------- | --------------------- | -------------------- |
| `draft`             | `sekdis_review`       | -                    |
| `sekdis_review`     | `evaluator_review`    | `draft`              |
| `evaluator_review`  | `ketua_review`        | `draft`              |
| `ketua_review`      | `pengendali_review`   | `evaluator_review`   |
| `pengendali_review` | `irban_review`        | `evaluator_review`   |
| `irban_review`      | `final`               | `evaluator_review`   |

---

## TODO List

### Step 1: Database Schema

- [x] Create migration file `db/20260206_lke_rekap_status_workflow.sql`
- [x] Add CHECK constraint for valid status values
- [x] Create index on `status_evaluasi` for faster queries
- [x] Add `id_sekdis` and `id_irban` columns
- [x] Add `status_label`, `previous_status`, `status_description` columns

### Step 2: Update Model Layer (`models/lke_rekap.go`)

- [x] Add status constants (draft, sekdis_review, evaluator_review, ketua_review, pengendali_review, irban_review, final)
- [x] Add validTransitions map with rejection rules
- [x] Add `IsValidTransition(current, new string) bool` method
- [x] Add `UpdateStatus(id int64, newStatus string) error` method with status tracking
- [x] Add `GetByStatus(status string) ([]LkeRekap, error)` method
- [ ] Add `GetPendingByReviewer(userID int64, status string) ([]LkeRekap, error)` method
- [x] Update `Create` method to set initial status and labels
- [x] Update `One` and `OneWithEvaluasi` methods
- [x] Add `ResetToDraft` admin function
- [x] Add `GetIDSekdisByStatusEvaluasi` helper method

### Step 3: Update Form Validation (`forms/lke_rekap.go`)

- [x] Update `CreateLkeRekapForm` - includes StatusEvaluasi
- [x] Update `UpdateLkeRekapForm` - StatusEvaluasi optional
- [ ] Add validation for status values only (in support of `validTransitions`)

### Step 4: Add Workflow Service (`services/workflow.go` - NEW)

- [ ] Create `LKEWorkflowService` struct
- [ ] Add `Transition(id int64, action string, userID int64) error` method
- [ ] Add permission check: userID must match reviewer ID for that status
- [ ] Add logging for each transition

### Step 5: Update Controller Layer (`controllers/lke_rekap.go`)

- [x] Add `SubmitToSekdis(id int64, userID int64) error` - `draft` → `sekdis_review`
- [x] Add `ApproveBySekdis(id int64, userID int64) error` - `sekdis_review` → `evaluator_review`
- [x] Add `RejectBySekdis(id int64, userID int64) error` - `sekdis_review` → `draft`
- [x] Add `ApproveByEvaluator(id int64, userID int64) error` - `evaluator_review` → `ketua_review`
- [x] Add `RejectByEvaluator(id int64, userID int64) error` - `evaluator_review` → `draft`
- [x] Add `ApproveByKetua(id int64, userID int64) error` - `ketua_review` → `pengendali_review`
- [x] Add `RejectByKetua(id int64, userID int64) error` - `ketua_review` → `evaluator_review`
- [x] Add `ApproveByPengendali(id int64, userID int64) error` - `pengendali_review` → `irban_review`
- [x] Add `RejectByPengendali(id int64, userID int64) error` - `pengendali_review` → `evaluator_review`
- [x] Add `ApproveByIrban(id int64, userID int64) error` - `irban_review` → `final`
- [x] Add `RejectByIrban(id int64, userID int64) error` - `irban_review` → `evaluator_review`
- [x] Add `ResetToDraft` admin endpoint

### Step 6: Add API Endpoints

- [x] `POST /v1/lke-rekap/:id/submit-sekdis` - Submit to SEKDIS review
- [x] `POST /v1/lke-rekap/:id/approve-sekdis` - SEKDIS approval → evaluator_review
- [x] `POST /v1/lke-rekap/:id/reject-sekdis` - SEKDIS rejection → draft
- [x] `POST /v1/lke-rekap/:id/approve-evaluator` - Evaluator approval → ketua_review
- [x] `POST /v1/lke-rekap/:id/reject-evaluator` - Evaluator rejection → draft
- [x] `POST /v1/lke-rekap/:id/approve-ketua` - KETUA approval → pengendali_review
- [x] `POST /v1/lke-rekap/:id/reject-ketua` - KETUA rejection → evaluator_review
- [x] `POST /v1/lke-rekap/:id/approve-pengendali` - PENGENDALI approval → irban_review
- [x] `POST /v1/lke-rekap/:id/reject-pengendali` - PENGENDALI rejection → evaluator_review
- [x] `POST /v1/lke-rekap/:id/approve-irban` - IRBAN approval → final
- [x] `POST /v1/lke-rekap/:id/reject-irban` - IRBAN rejection → evaluator_review
- [x] `POST /v1/lke-rekap/:id/reset` - Reset to draft (admin)

### Step 7: Permission Checking

- [ ] Add helper method `CanUserReview(userID int64, lkeRekapID int64) bool`
- [ ] Check: userID must match `id_sekdis` when status is `sekdis_review`
- [ ] Check: userID must match `id_evaluator` when status is `evaluator_review`
- [ ] Check: userID must match `id_ketua` when status is `ketua_review`
- [ ] Check: userID must match `id_pengendali` when status is `pengendali_review`
- [ ] Check: userID must match `id_irban` when status is `irban_review`
- [ ] Return 403 Forbidden if user is not authorized

**Note:** Permission checking is currently BYPASSED for testing purposes (see line 158)

### Step 8: Response Helper

- [x] Create `StatusTransitionResponse` struct
- [x] Include previous_status, new_status, message, status_label
- [x] Add `GetStatusLabel` helper function
- [x] Add `GetStatusDescription` and `GetForwardDescription` functions

### Step 9: Testing

- [ ] Unit tests for `IsValidTransition`
- [ ] Unit tests for `UpdateStatus`
- [ ] Unit tests for permission checks
- [ ] Integration tests for each endpoint
- [ ] Test invalid transitions return errors

### Step 10: Documentation

- [ ] Update Swagger/OpenAPI docs
- [ ] Add status workflow diagram
- [ ] Document API endpoints with examples

---

## File Changes Summary

### New Files

- `db/20260206_lke_rekap_status_workflow.sql` - Migration script ✓

### Modified Files

- `models/lke_rekap.go` - Add workflow methods ✓
- `forms/lke_rekap.go` - Update validation ✓
- `controllers/lke_rekap.go` - Add new endpoints ✓
- `main.go` - Register new routes ✓
- `docs/docs.go`, `docs/swagger.json`, `docs/swagger.yaml` - Update API docs

### Remaining Tasks

- Implement permission checking (Step 7)
- Add unit tests (Step 9)
- Update Swagger documentation (Step 10)

---

## API Endpoints Quick Reference

| Endpoint                                     | Status Change                            | Method | Controller Method           |
| -------------------------------------------- | ---------------------------------------- | ------ | --------------------------- |
| **Submit/Approval Endpoints**                |                                          |        |                             |
| `POST /v1/lke-rekap/:id/submit-sekdis`       | `draft` → `sekdis_review`                | POST   | SubmitToSekdis              |
| `POST /v1/lke-rekap/:id/approve-sekdis`      | `sekdis_review` → `evaluator_review`     | POST   | ApproveBySekdis             |
| `POST /v1/lke-rekap/:id/reject-sekdis`       | `sekdis_review` → `draft`                | POST   | RejectBySekdis              |
| `POST /v1/lke-rekap/:id/approve-evaluator`   | `evaluator_review` → `ketua_review`      | POST   | ApproveByEvaluator          |
| `POST /v1/lke-rekap/:id/reject-evaluator`    | `evaluator_review` → `draft`             | POST   | RejectByEvaluator           |
| `POST /v1/lke-rekap/:id/approve-ketua`       | `ketua_review` → `pengendali_review`     | POST   | ApproveByKetua              |
| `POST /v1/lke-rekap/:id/reject-ketua`        | `ketua_review` → `evaluator_review`      | POST   | RejectByKetua               |
| `POST /v1/lke-rekap/:id/approve-pengendali`  | `pengendali_review` → `irban_review`     | POST   | ApproveByPengendali         |
| `POST /v1/lke-rekap/:id/reject-pengendali`   | `pengendali_review` → `evaluator_review` | POST   | RejectByPengendali          |
| `POST /v1/lke-rekap/:id/approve-irban`       | `irban_review` → `final`                 | POST   | ApproveByIrban              |
| `POST /v1/lke-rekap/:id/reject-irban`        | `irban_review` → `evaluator_review`      | POST   | RejectByIrban               |
| **Admin Endpoints**                          |                                          |        |                             |
| `POST /v1/lke-rekap/:id/reset`               | any → `draft`                            | POST   | ResetToDraft                |
| **Query Endpoints**                          |                                          |        |                             |
| `GET /v1/lke-rekaps/sekdis-by-status/:tahun` | Get grouped ID Sekdis by status          | GET    | GetIDSekdisByStatusEvaluasi |

## Permission Status

**⚠️ Permission checking is BYPASSED for testing purposes**

All workflow endpoints currently accept any authenticated user. To enable proper authorization:

1. Implement Step 7 (Permission Checking)
2. Each endpoint should verify:
   - `submit-sekdis`: creator only (user_id matches)
   - `approve-sekdis`/`reject-sekdis`: user_id matches `id_sekdis`
   - `approve-evaluator`/`reject-evaluator`: user_id matches `id_evaluator`
   - `approve-ketua`/`reject-ketua`: user_id matches `id_ketua`
   - `approve-pengendali`/`reject-pengendali`: user_id matches `id_pengendali`
   - `approve-irban`/`reject-irban`: user_id matches `id_irban`
   - `reset`: admin only

## Implementation Order

1. Database migration (status constants + constraint)
2. Model layer (constants + validation methods)
3. Workflow service (transition logic + permissions)
4. Controller endpoints
5. Testing
6. Documentation
