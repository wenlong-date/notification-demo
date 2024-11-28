package main

import (
	"notification/helper"
	"notification/notification"
	"time"
)

func main() {
	// save handler
	saveTask()

	// get notification list handler
	notification.GetNotificationList(1, 1, time.Now().Unix(), 10)
}

var demoNotifications = []struct {
	NotificationType notification.NotificationType
	Notification     notification.NotificationInterface
}{
	{
		NotificationType: notification.NotificationTypeTask,
		Notification: &notification.TaskNotification{
			CommonNotification: notification.NewCommonNotification(1, 2, 3),
			TaskNotificationDetail: notification.TaskNotificationDetail{
				Creator:      1,
				CreatorName:  "creator",
				Director:     2,
				DirectorName: "director",
				SubType:      notification.NotificationTypeTaskMovieShoot,
				TaskID:       1,
				TaskName:     "task name",
			},
		},
	},
	{
		NotificationType: notification.NotificationTypeTask,
		Notification: &notification.TaskNotification{
			CommonNotification: notification.NewCommonNotification(1, 2, 3),
			TaskNotificationDetail: notification.TaskNotificationDetail{
				Creator:      1,
				CreatorName:  "creator",
				Director:     3,
				DirectorName: "director3",
				SubType:      notification.NotificationTypeTaskMovieMontage,
				TaskID:       1,
				TaskName:     "task name1",
			},
		},
	},
	{
		NotificationType: notification.NotificationTypeBroadcastSystemUpdate,
		Notification: &notification.BroadcastSystemUpdateNotification{
			BroadcastNotificationDetail: notification.BroadcastNotificationDetail{
				Msg: "broadcast system update",
			},
			CommonBroadcastNotification: notification.DefaultCommonBroadcastNotification,
		},
	},
	{
		NotificationType: notification.NotificationTypePlan,
		Notification: &notification.StatusChangeNotification{
			CommonNotification: notification.NewCommonNotification(1, 2, 3),
			StatusChangeNotificationDetail: notification.StatusChangeNotificationDetail{
				FromStatus: "待审核",
				PlanID:     1,
				PlanName:   "plan name",
				ToStatus:   "已审核上架",
			},
		},
	},
}

func saveTask() {
	for _, n := range demoNotifications {
		err := notification.DispatcherNotification(n.Notification)
		if err != nil {
			helper.Dump("err is ", err)
		}
		helper.Dump("dispatcher notification success")
	}
}
