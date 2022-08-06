package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NotificationToUser struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserId         int                `bson:"userId,omitempty"`
	NotificationId string             `bson:"notificationId,omitempty"`
	Seen           int                `bson:"seen,omitempty"`
	CreatedAt      *time.Time         `bson:"createdAt,omitempty"`
	UpdatedAt      *time.Time         `bson:"updatedAt,omitempty"`
}

func (NotificationToUser) CollectionName() string {
	return "notification_to_user"
}
