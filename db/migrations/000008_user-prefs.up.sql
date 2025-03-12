BEGIN;

CREATE TABLE user_preferences (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    pref_value TEXT NOT NULL,
    dont_disturb_start TIME NOT NULL,
    dont_disturb_end TIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

COMMIT;