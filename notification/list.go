package notification

import (
	"notification/helper"
	"time"
)

type UserInfo struct {
	ID     string // hash id
	Name   string
	Avatar string
}

type NotificationItem struct {
	ID               int
	NotificationType NotificationType
	Msg              string
	DetailJson       string `json:"-"`
	Detail           NotificationDetailInterface
	FromUserID       int
	FromUserInfo     UserInfo
	ToUserID         int
	ToUserInfo       UserInfo
	CompanyID        int
	IsRead           int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ListResp struct {
	Records []NotificationItem
	Count   int64
}

var itemNow = time.Now()

var items = []NotificationItem{
	{
		ID:               1,
		NotificationType: NotificationTypeBroadcastSystemUpdate,
		Msg:              "系统更新 xxx ",
		DetailJson:       "系统更新 xxx",
		FromUserID:       0,
		ToUserID:         0,
		CompanyID:        1,
		IsRead:           0,
		CreatedAt:        itemNow,
		UpdatedAt:        itemNow,
	},
	{
		ID:               1,
		NotificationType: NotificationTypeTask,
		Msg:              "剪片任务1 xxx ",
		DetailJson:       `{"SubType":202,"TaskID":1,"TaskName":"task name","Creator":1,"CreatorName":"creator","Director":2,"DirectorName":"director"}`,
		FromUserID:       1,
		ToUserID:         2,
		CompanyID:        1,
		IsRead:           0,
		CreatedAt:        itemNow,
		UpdatedAt:        itemNow,
	},
	{
		ID:               2,
		NotificationType: NotificationTypeTask,
		Msg:              "剪片任务2 xxx ",
		DetailJson:       `{"SubType":202,"TaskID":1,"TaskName":"task name","Creator":1,"CreatorName":"creator","Director":2,"DirectorName":"director"}`,
		FromUserID:       1,
		ToUserID:         2,
		CompanyID:        1,
		IsRead:           0,
		CreatedAt:        itemNow,
		UpdatedAt:        itemNow,
	},
	{
		ID:               3,
		NotificationType: NotificationTypePlan,
		Msg:              "《广告计划1》 状态变更",
		DetailJson:       `{"PlanID":1,"PlanName":"广告计划1","FromStatus":"待审核","ToStatus":"已审核上架"}`,
		FromUserID:       1,
		ToUserID:         2,
		CompanyID:        1,
		IsRead:           0,
		CreatedAt:        itemNow,
		UpdatedAt:        itemNow,
	},
}

func GetNotificationList(companyID, userID int, fromSentTime int64, limit int) (ListResp, error) {
	// 从db获取个人消息列表
	// 获取时间范围，获取范围内系统消息
	// 组装列表 成 []NotificationItem
	// 获取所有的用户ID，并找到 ID和用户详情的数据map map[userID]struct{ID int, Name string, Avatar string}

	for k, item := range items {
		// 填充 FromUserInfo ToUserInfo
		detail, err := TransferNotificationDetailFromJson(item.NotificationType, item.DetailJson)
		if err != nil {
			continue
		}
		item.Detail = detail
		items[k] = item
	}

	helper.DumpJson("GetNotificationList", items)

	return ListResp{
		Count:   2,
		Records: items,
	}, nil
}
