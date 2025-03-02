BEGIN;

-- Drop the existing primary key constraint (if it exists)
ALTER TABLE notif_states DROP CONSTRAINT IF EXISTS notif_states_pkey;

-- Create a new sequence for the id column
CREATE SEQUENCE notif_states_id_seq;

-- Update the id column to use the sequence and set it as the primary key
ALTER TABLE notif_states
    ALTER COLUMN id SET DEFAULT nextval('notif_states_id_seq'),
    ALTER COLUMN id SET NOT NULL;

-- Set the sequence's current value to the maximum id in the table (if data exists)
SELECT setval('notif_states_id_seq', COALESCE((SELECT MAX(id) FROM notif_states), 1), false);

-- Add the primary key constraint
ALTER TABLE notif_states ADD PRIMARY KEY (id);

COMMIT;