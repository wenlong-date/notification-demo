package notification

import (
	"notification/helper"
)

func DispatcherNotification(n NotificationInterface) error {
	// 也可以通过其他方式判断是否是系统消息，例如指定配置数组
	if n.GetNotificationType() < 200 {
		return handlerBroadcastNotification(n)
	}

	return handlerIndividualNotification(n)
}

// 也可以考虑做成 handler func map 的形式
func handlerBroadcastNotification(n NotificationInterface) error {
	// TODO 全局广播的消息处理

	helper.DumpJson("handlerBroadcastNotification", TableBroadcastNotification{
		Msg: n.GetMessage(),
	})

	// save to db
	return nil
}

func handlerIndividualNotification(n NotificationInterface) error {
	// TODO 个人消息处理
	detailJson, err := n.GetSaveToDBJson()
	if err != nil {
		return err
	}
	helper.DumpJson("handlerIndividualNotification", TableIndividualNotification{
		SenderID:         n.GetFromUserID(),
		ReceiverID:       n.GetToUserID(),
		CompanyID:        n.GetCompanyID(),
		NotificationType: n.GetNotificationType(),
		Msg:              n.GetMessage(),
		Detail:           detailJson,
	})

	// save to db
	return nil
}
