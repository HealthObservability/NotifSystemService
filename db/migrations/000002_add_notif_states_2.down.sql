BEGIN;

DROP SEQUENCE IF EXISTS notif_states_id_seq;
DROP SEQUENCE IF EXISTS notif_info_id_seq;

DROP TYPE IF EXISTS notif_repeat_interval;

DROP TABLE IF EXISTS notif_states CASCADE;

DROP TABLE IF EXISTS notif_info CASCADE;


COMMIT;