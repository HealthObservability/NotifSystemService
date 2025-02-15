# NotifSystemService

Is a service that writes a user notification settings to database,
fetch it and creates notifications to another db...

## Diagram

check for [DRAW.IO STRUCTURE FILE](notif.drawio)

## TODO
- [ ] update notifications on user preferences changes
- [ ] fetch notifications status table every N seconds
- [ ] send notification after fetch to sender service by http


- [ ] Try Amazon Cognito to maintain users info
- [ ] Try Amazon DynamoDB for User preferences

_P.S.
reference: https://www.notificationapi.com/_