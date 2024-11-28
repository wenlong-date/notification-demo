package notification

type BroadcastSystemUpdateNotification struct {
	CommonBroadcastNotification
	BroadcastNotificationDetail
}

type BroadcastNotificationDetail struct {
	Msg string
}

func (b *BroadcastNotificationDetail) GetMessage() string {
	return b.Msg
}
func (b *BroadcastNotificationDetail) GetNotificationType() NotificationType {
	return NotificationTypeBroadcastSystemUpdate
}
func (b *BroadcastNotificationDetail) GetSaveToDBJson() (string, error) {
	return b.Msg, nil
}

func TransferBroadcastSystemUpdateNotificationDetail(s string) (NotificationDetailInterface, error) {
	return &BroadcastNotificationDetail{Msg: s}, nil
}

func init() {
	AddNotificationDetailHandlerFunc(NotificationTypeBroadcastSystemUpdate, TransferBroadcastSystemUpdateNotificationDetail)
}
