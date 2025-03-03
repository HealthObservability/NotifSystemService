BEGIN;

CREATE SEQUENCE IF NOT EXISTS notif_states_id_seq;

CREATE TABLE IF NOT EXISTS notif_states
(
    id BIGINT PRIMARY KEY DEFAULT nextval('notif_states_id_seq'),
    notif_id BIGINT NOT NULL,
    send_due TIMESTAMP WITHOUT TIME ZONE,
    creation_timestamp TIMESTAMP WITHOUT TIME ZONE,
    status VARCHAR(25)
);

COMMIT;
