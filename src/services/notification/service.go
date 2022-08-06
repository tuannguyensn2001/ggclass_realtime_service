package notification

import (
	"context"
	"ggclass_log_service/src/logger"
	"ggclass_log_service/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	CreateNotificationToUser(ctx context.Context, list []models.NotificationToUser) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input createNotificationInput) (string, error) {
	now := time.Now()
	notification := models.Notification{
		OwnerAvatar: input.OwnerAvatar,
		OwnerName:   input.OwnerName,
		CreatedBy:   input.CreatedBy,
		HtmlContent: input.HtmlContent,
		ClassId:     input.ClassId,
		Content:     input.Content,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		ID:          primitive.NewObjectID(),
	}

	err := s.repository.Create(ctx, &notification)
	if err != nil {
		return "", err
	}

	return notification.ID.Hex(), nil
}

func (s *service) NotifyToUser(ctx context.Context, notificationId string, users []int) error {
	list := make([]models.NotificationToUser, len(users))

	logger.Sugar().Info(users)

	for index, item := range users {
		now := time.Now()
		list[index] = models.NotificationToUser{
			NotificationId: notificationId,
			UserId:         item,
			Seen:           0,
			CreatedAt:      &now,
			UpdatedAt:      &now,
			ID:             primitive.NewObjectID(),
		}
	}

	logger.Sugar().Info(list)
	return s.repository.CreateNotificationToUser(ctx, list)

}
