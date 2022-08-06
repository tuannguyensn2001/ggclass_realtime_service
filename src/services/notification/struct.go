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
