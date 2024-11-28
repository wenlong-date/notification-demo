package notification

import "encoding/json"

type TaskNotification struct {
	CommonNotification
	TaskNotificationDetail
}

type TaskNotificationDetail struct {
	SubType      NotificationSubType
	TaskID       int
	TaskName     string
	Creator      int
	CreatorName  string
	Director     int
	DirectorName string
}

func (t *TaskNotificationDetail) GetMessage() string {
	return t.TaskName
}

func (t *TaskNotificationDetail) GetNotificationType() NotificationType {
	return NotificationTypeTask
}

func (t *TaskNotificationDetail) GetSaveToDBJson() (string, error) {
	s, err := json.Marshal(*t)
	return string(s), err
}

func TransferTaskNotificationDetail(s string) (NotificationDetailInterface, error) {
	var t TaskNotificationDetail
	err := json.Unmarshal([]byte(s), &t)
	return &t, err
}

func init() {
	AddNotificationDetailHandlerFunc(NotificationTypeTask, TransferTaskNotificationDetail)
}
