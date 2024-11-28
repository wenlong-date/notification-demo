package notification

type NotificationType int
type NotificationSubType int

const (
	// 系统通知
	NotificationTypeBroadcast             NotificationType = 100
	NotificationTypeBroadcastSystemUpdate NotificationType = 101

	// 任务通知
	NotificationTypeTask             NotificationType    = 200
	NotificationTypeTaskMovieMontage NotificationSubType = 201
	NotificationTypeTaskMovieShoot   NotificationSubType = 202
	NotificationTypeTaskMovieScirpt  NotificationSubType = 203
	NotificationTypeTaskMovieDoc     NotificationSubType = 204
	NotificationTypeTaskMovieAudit   NotificationSubType = 205

	// 爆款通知
	NotificationTypeHot NotificationType = 300

	// Plan
	NotificationTypePlan NotificationType = 400
)
