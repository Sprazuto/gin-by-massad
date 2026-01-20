-- Migration: Add manual_status_override field to vehicle_asset table
-- This field tracks when status was manually changed to prevent auto-update overrides

ALTER TABLE vehicle_asset
ADD COLUMN manual_status_override BOOLEAN DEFAULT FALSE;

-- Add comment for documentation
COMMENT ON COLUMN vehicle_asset.manual_status_override IS 'Indicates if status was manually set and should not be auto-updated';