BEGIN;
ALTER TABLE notif_info RENAME COLUMN last_notif TO last_scheduled_time;
COMMIT;