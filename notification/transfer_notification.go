package notification

import "errors"

type NotificationDetailHandlerFunc func(s string) (NotificationDetailInterface, error)

var notificationDetailHandlerFuncMap = map[NotificationType]NotificationDetailHandlerFunc{}

func AddNotificationDetailHandlerFunc(t NotificationType, f NotificationDetailHandlerFunc) {
	notificationDetailHandlerFuncMap[t] = f
}

func GetNotificationDetailHandlerFunc(t NotificationType) (NotificationDetailHandlerFunc, bool) {
	f, ok := notificationDetailHandlerFuncMap[t]
	return f, ok
}

func TransferNotificationDetailFromJson(t NotificationType, s string) (NotificationDetailInterface, error) {
	var (
		nd NotificationDetailInterface
	)

	f, ok := GetNotificationDetailHandlerFunc(t)
	if !ok {
		return nd, errors.New("Notification type not found")
	}

	return f(s)
}
