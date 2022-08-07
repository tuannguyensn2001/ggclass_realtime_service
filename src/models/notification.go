package models

import (
	"ggclass_log_service/src/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID          primitive.ObjectID     `bson:"_id"`
	OwnerName   string                 `bson:"ownerName,omitempty"`
	OwnerAvatar string                 `bson:"ownerAvatar,omitempty"`
	HtmlContent string                 `bson:"htmlContent,omitempty"`
	ClassId     int                    `bson:"classId,omitempty"`
	CreatedBy   int                    `bson:"createdBy,omitempty"`
	Content     string                 `bson:"content,omitempty"`
	CreatedAt   *time.Time             `bson:"createdAt,omitempty"`
	UpdatedAt   *time.Time             `bson:"updatedAt,omitempty"`
	Type        enums.NotificationType `bson:"type,omitempty"`
	Seen        int
}

func (Notification) CollectionName() string {
	return "notifications"
}
