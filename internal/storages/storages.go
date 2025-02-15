package storages

type NotificationStates interface {
	UpdateNotifications()
	GetNotifications()
}

type Notifications interface {
	AddNotification()
	AddNotifications()
}

type Storage struct {
	NotificationStates
	Notifications
}
