-- Migration: LKE Rekap Status Workflow + Approval Columns
-- Version: 1.7.0 (combined from 20260203 + 20260206)
-- Date: 2026-02-06
-- Description: Adds status_label field, CHECK constraint for status_evaluasi, and approval columns
-- Combined from: 20260203_add_approval_columns.sql + 20260206_lke_rekap_status_workflow.sql

BEGIN;

-- ============================================
-- PART 1: Add approval columns (from 20260203)
-- ============================================

-- Add the id_sekdis column if it doesn't exist (for Sekretaris Daerah approval)
ALTER TABLE lke_rekap ADD COLUMN IF NOT EXISTS id_sekdis INTEGER;

-- Add the id_irban column if it doesn't exist (for IRBAN approval)
ALTER TABLE lke_rekap ADD COLUMN IF NOT EXISTS id_irban INTEGER;

-- Create index for id_sekdis
CREATE INDEX IF NOT EXISTS idx_lke_rekap_id_sekdis ON public.lke_rekap (id_sekdis);

-- Create index for id_irban
CREATE INDEX IF NOT EXISTS idx_lke_rekap_id_irban ON public.lke_rekap (id_irban);

-- ============================================
-- PART 2: Add status_label, previous_status, and status_description fields
-- ============================================

-- Add status_label column if it doesn't exist
ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS status_label character varying;

-- Add previous_status column to track where draft came from
-- NULL = fresh draft, otherwise = the status before transition to draft
ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS previous_status character varying;

-- Add status_description for human-readable rejection/return notes
-- Example: "Draft dikembalikan oleh Evaluator untuk perbaikan"
ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS status_description character varying;

-- Update legacy status values to new standardized values
UPDATE public.lke_rekap
SET
    status_evaluasi = CASE status_evaluasi
        -- Legacy values mapping
        WHEN 'Belum Dievaluasi' THEN 'draft'
        WHEN 'Belum Diverifikasi' THEN 'evaluator_review'
        WHEN 'Menunggu Tanggapan Sekdis' THEN 'sekdis_review'
        WHEN 'Perlu Revisi' THEN 'draft' -- Returned for revision, goes back to draft
        WHEN 'Direset' THEN 'draft' -- Reset, goes back to draft
        WHEN 'Sudah Dievaluasi' THEN 'final'
        -- Standard values (no change)
        WHEN 'draft' THEN 'draft'
        WHEN 'sekdis_review' THEN 'sekdis_review'
        WHEN 'evaluator_review' THEN 'evaluator_review'
        WHEN 'ketua_review' THEN 'ketua_review'
        WHEN 'pengendali_review' THEN 'pengendali_review'
        WHEN 'irban_review' THEN 'irban_review'
        WHEN 'final' THEN 'final'
        ELSE 'draft' -- Default for any unknown values
    END;

-- Initialize status_label for all records based on new status_evaluasi
UPDATE public.lke_rekap
SET
    status_label = CASE status_evaluasi
        WHEN 'draft' THEN 'Belum Dievaluasi'
        WHEN 'sekdis_review' THEN 'Review Sekdis'
        WHEN 'evaluator_review' THEN 'Review Evaluator'
        WHEN 'ketua_review' THEN 'Review Ketua'
        WHEN 'pengendali_review' THEN 'Review Pengendali'
        WHEN 'irban_review' THEN 'Review IRBAN'
        WHEN 'final' THEN 'Sudah Dievaluasi'
        ELSE 'Belum Dievaluasi'
    END;

-- Initialize status_description for all records based on status_evaluasi
-- This provides initial human-readable descriptions for existing records
UPDATE public.lke_rekap
SET
    status_description = CASE status_evaluasi
        WHEN 'draft' THEN 'Belum dievaluasi'
        WHEN 'sekdis_review' THEN 'Diteruskan ke Sekretaris untuk review'
        WHEN 'evaluator_review' THEN 'Diteruskan ke Evaluator oleh Sekretaris'
        WHEN 'ketua_review' THEN 'Diteruskan ke Ketua Tim oleh Evaluator'
        WHEN 'pengendali_review' THEN 'Diteruskan ke Pengendali teknis oleh Ketua Tim'
        WHEN 'irban_review' THEN 'Diteruskan ke Inspektur Pembantu oleh Pengendali Teknis'
        WHEN 'final' THEN 'Sudah dievaluasi'
        ELSE NULL
    END;

-- Set NOT NULL constraint on status_label
ALTER TABLE lke_rekap ALTER COLUMN status_label SET NOT NULL;

-- Add CHECK constraint for valid status_evaluasi values
ALTER TABLE lke_rekap DROP CONSTRAINT IF EXISTS chk_status_evaluasi;

ALTER TABLE lke_rekap
ADD CONSTRAINT chk_status_evaluasi CHECK (
    status_evaluasi IN (
        'draft',
        'sekdis_review',
        'evaluator_review',
        'ketua_review',
        'pengendali_review',
        'irban_review',
        'final'
    )
);

-- Create index for faster status_evaluasi queries
CREATE INDEX IF NOT EXISTS idx_lke_rekap_status_evaluasi ON public.lke_rekap (status_evaluasi);

-- Create index for faster status_label queries
CREATE INDEX IF NOT EXISTS idx_lke_rekap_status_label ON public.lke_rekap (status_label);

-- Create index for previous_status queries
CREATE INDEX IF NOT EXISTS idx_lke_rekap_previous_status ON public.lke_rekap (previous_status);

COMMIT;

-- ============================================
-- Verification (runs outside transaction)
-- ============================================

-- Verify the constraint is working
-- This will fail if there are invalid status values
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM public.lke_rekap
        WHERE status_evaluasi NOT IN (
            'draft',
            'sekdis_review',
            'evaluator_review',
            'ketua_review',
            'pengendali_review',
            'irban_review',
            'final'
        )
    ) THEN
        RAISE EXCEPTION 'Found records with invalid status_evaluasi values';
    END IF;
    RAISE NOTICE 'All status_evaluasi values are valid';
END $$;

-- ============================================
-- Comments for documentation
-- ============================================
COMMENT ON CONSTRAINT chk_status_evaluasi ON public.lke_rekap IS 'Validates that status_evaluasi is one of the allowed workflow states';

COMMENT ON COLUMN public.lke_rekap.status_label IS 'Human-readable Indonesian label for status_evaluasi (e.g., "Belum Dievaluasi", "Review Sekdis")';

COMMENT ON COLUMN public.lke_rekap.previous_status IS 'Tracks the previous status before transition to draft (NULL for fresh drafts)';

COMMENT ON COLUMN public.lke_rekap.status_description IS 'Human-readable description for draft status (e.g., "Draft dikembalikan oleh Evaluator untuk perbaikan")';

COMMENT ON COLUMN lke_rekap.id_sekdis IS 'ID of Sekretaris Daerah who approved this LKE evaluation';

COMMENT ON COLUMN lke_rekap.id_irban IS 'ID of IRBAN (Inspektorat/BPK) who reviewed/approved this LKE evaluation';

-- Show final status distribution
SELECT
    status_evaluasi,
    status_label,
    status_description,
    COUNT(*) as count
FROM public.lke_rekap
GROUP BY
    status_evaluasi,
    status_label,
    status_description
ORDER BY status_evaluasi;