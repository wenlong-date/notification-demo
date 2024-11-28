package notification

var DefaultCommonBroadcastNotification CommonBroadcastNotification

func init() {
	DefaultCommonBroadcastNotification = CommonBroadcastNotification{}
}

type CommonBroadcastNotification struct {
}

func (cn *CommonBroadcastNotification) GetFromUserID() int {
	return 0
}

func (cn *CommonBroadcastNotification) GetToUserID() int {
	return 0
}

func (cn *CommonBroadcastNotification) GetCompanyID() int {
	return 0
}
