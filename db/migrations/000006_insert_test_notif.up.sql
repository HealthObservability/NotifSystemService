-- INSERT INTO notif_states (notif_id, send_due, creation_timestamp, status)
-- VALUES (1, '3/2/2025, 10:30:27 AM', '3/2/2025, 2:30:00 AM',)

WITH inserted_notif AS (
    INSERT INTO notif_info(creation_timestamp, id_to_send, text, repeat_interval, since_time)
    VALUES('3/2/2025, 10:30:27 AM', 2, 'notification text', 'minute', '3/2/2025, 10:30:27 AM')
    RETURNING id
)
INSERT INTO notif_states(notif_id, send_due, creation_timestamp, status)
SELECT id, '3/2/2025, 10:31:27 AM', '3/2/2025, 10:30:27 AM', 'not_sent'
FROM inserted_notif;
