BEGIN;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'notif_repeat_interval') THEN
            CREATE TYPE notif_repeat_interval AS ENUM('no-repeat', 'minute', 'half-hour', 'hour', 'day', 'week', 'month');
        END IF;
END $$;


CREATE TABLE IF NOT EXISTS notif_info
(
    id SERIAL PRIMARY KEY,
    creation_timestamp timestamp without time zone NOT NULL DEFAULT now(),
    id_to_send BIGINT NOT NULL,
    text TEXT NOT NULL,
    repeat_interval notif_repeat_interval NOT NULL,
    last_notif  timestamp without time zone,
    since_time timestamp without time zone NOT NULL
);

ALTER TABLE notif_states
    ADD FOREIGN KEY (notif_id) REFERENCES notif_info (id);

COMMIT;