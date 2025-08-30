-- SQL script to fix the duplicate key issue in vehicle_asset table
-- Run these commands in your PostgreSQL database

-- Step 1: Check the current maximum ID in the vehicle_asset table
SELECT MAX(id) as max_id FROM vehicle_asset;

-- Step 2: Check the current sequence value
SELECT nextval('vehicle_asset_id_seq') as current_sequence_value;

-- Step 3: If the sequence value is less than or equal to max_id, reset the sequence
-- Replace [max_id_from_step1] with the actual max_id value from step 1
-- SELECT setval('vehicle_asset_id_seq', [max_id_from_step1]);

-- Alternative: Reset sequence to MAX(id) + 1 to be safe
SELECT setval(
        'vehicle_asset_id_seq', (
            SELECT MAX(id)
            FROM vehicle_asset
        ) + 1
    );

-- Step 4: Verify the sequence is now correct
SELECT nextval('vehicle_asset_id_seq') as new_sequence_value;