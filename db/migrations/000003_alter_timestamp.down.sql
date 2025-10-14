-- Step 1: Add the original timestamp column (TIMESTAMPTZ)
ALTER TABLE vehicle_location ADD COLUMN timestampZ TIMESTAMPTZ;

-- Step 2: Populate the new timestamp column with the BIGINT values (epoch time)
UPDATE vehicle_location 
SET timestampZ = TO_TIMESTAMP(timestamp_bigint);

-- Step 3: Drop the timestamp_bigint column (which was used to store BIGINT)
ALTER TABLE vehicle_location DROP COLUMN timestamp_bigint;

-- Step 4: Rename the new column to 'timestamp'
ALTER TABLE vehicle_location RENAME COLUMN timestampZ TO timestamp;