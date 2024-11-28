package notification

import "encoding/json"

type StatusChangeNotification struct {
	CommonNotification
	StatusChangeNotificationDetail
}

type StatusChangeNotificationDetail struct {
	PlanID     int
	PlanName   string
	FromStatus string
	ToStatus   string
}

func (t *StatusChangeNotificationDetail) GetMessage() string {
	return t.PlanName + "计划状态由" + t.FromStatus + "变为" + t.ToStatus
}

func (t *StatusChangeNotificationDetail) GetNotificationType() NotificationType {
	return NotificationTypePlan
}

func (t *StatusChangeNotificationDetail) GetSaveToDBJson() (string, error) {
	s, err := json.Marshal(*t)
	return string(s), err
}

func TransferStatusChangeNotificationDetail(s string) (NotificationDetailInterface, error) {
	var t StatusChangeNotificationDetail
	err := json.Unmarshal([]byte(s), &t)
	return &t, err
}

func init() {
	AddNotificationDetailHandlerFunc(NotificationTypePlan, TransferStatusChangeNotificationDetail)
}
