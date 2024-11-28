package notification

type NotificationInterface interface {
	CommonNotificationInterface
	NotificationDetailInterface
}

type NotificationDetailInterface interface {
	GetMessage() string
	GetNotificationType() NotificationType
	GetSaveToDBJson() (string, error)
}
