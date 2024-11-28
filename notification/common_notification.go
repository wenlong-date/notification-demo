package notification

type CommonNotificationInterface interface {
	GetFromUserID() int
	GetToUserID() int
	GetCompanyID() int
}

type CommonNotification struct {
	FromUserID int
	ToUserID   int
	CompanyID  int
}

func NewCommonNotification(fromUserID, toUserID, companyID int) CommonNotification {
	return CommonNotification{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		CompanyID:  companyID,
	}
}

func (cn *CommonNotification) GetFromUserID() int {
	return cn.FromUserID
}

func (cn *CommonNotification) GetToUserID() int {
	return cn.ToUserID
}

func (cn *CommonNotification) GetCompanyID() int {
	return cn.CompanyID
}
