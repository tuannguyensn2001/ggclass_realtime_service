package notification

type createNotificationInput struct {
	OwnerName   string
	OwnerAvatar string
	CreatedBy   int
	HtmlContent string
	ClassId     int
	Content     string
}

type notifyToUser struct {
	NotificationId string `json:"id"`
	Users          []int  `json:"users"`
}

type setSeenInput struct {
	UserId         int    `json:"userId"`
	NotificationId string `json:"notificationId"`
}
