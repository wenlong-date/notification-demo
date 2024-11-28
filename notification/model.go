package notification

import "time"

// TableBroadcastNotification 系统通知
type TableBroadcastNotification struct {
	ID        int
	Msg       string
	SentAt    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableIndividualNotification 个人通知
type TableIndividualNotification struct {
	ID               int
	SenderID         int
	ReceiverID       int
	CompanyID        int
	NotificationType NotificationType
	Msg              string
	Detail           string
	IsRead           int
	IsDeleted        int
	SentAt           int64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
