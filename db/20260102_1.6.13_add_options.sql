-- Migration: Add options column to lke_komponen table
-- Version: 1.6.13
-- Date: 2026-01-02
-- Description: Adds a JSONB column 'options' to store selectable form options per row
-- Safe to run on active database

BEGIN;

-- Add the options column if it doesn't exist
ALTER TABLE lke_komponen ADD COLUMN IF NOT EXISTS options JSONB;

COMMIT;

-- ["Ya", "Tidak"]
-- ["Sudah", "Sebagian", "Belum"]
-- ["Ya", "Sebagian", "Tidak"]