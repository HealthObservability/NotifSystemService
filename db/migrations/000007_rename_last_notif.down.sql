BEGIN;
ALTER TABLE notif_info RENAME COLUMN last_scheduled_time TO last_notif;
COMMIT;