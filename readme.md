# NotifSystemService

Is a service that writes a user notification settings to database,
fetch it and creates notifications to another db...

## Diagram

check for [DRAW.IO STRUCTURE FILE](notif.drawio)

## Expected stack
- Amazon DynamoDB
- Amazon Cognito
- RabbitMQ as producer to senders and as listener for notif updates
- Postgres as  

## TODO
- [x] fetch notifications status table every N seconds
- [x] send notification after fetch to sender service by http
- [x] schedule notification time
- [x] schedule repeated notification without spamming on server failure


- [ ] User preferences
- [ ] schedule notifications based on user preferences

_P.S.
reference: https://www.notificationapi.com/_