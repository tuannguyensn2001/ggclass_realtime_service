package notification

type createNotificationInput struct {
	OwnerName   string
	OwnerAvatar string
	CreatedBy   int
	HtmlContent string
	ClassId     int
	Content     string
}
