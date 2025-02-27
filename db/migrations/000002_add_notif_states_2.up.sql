BEGIN;

CREATE TABLE IF NOT EXISTS notif_states
(
    id bigint PRIMARY KEY,
    notif_id bigint NOT NULL,
    send_due timestamp without time zone,
    creation_timestamp timestamp without time zone,
    status varchar(25)
);

COMMIT;