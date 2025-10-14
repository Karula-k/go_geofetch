-- Step 1: Add the new column to store the timestamp as BIGINT
ALTER TABLE vehicle_location ADD COLUMN timestamp_bigint BIGINT;

-- Step 2: Populate the new column with the values from the existing timestamp column
UPDATE vehicle_location 
SET timestamp_bigint = EXTRACT(EPOCH FROM timestamp)::BIGINT;

-- Step 3: Drop the old timestamp column
ALTER TABLE vehicle_location DROP COLUMN timestamp;

-- Step 4: Rename the new column to 'timestamp'
ALTER TABLE vehicle_location RENAME COLUMN timestamp_bigint TO timestamp;

-- Step 5: add not null constraint
ALTER TABLE vehicle_location ALTER COLUMN timestamp SET NOT NULL;
