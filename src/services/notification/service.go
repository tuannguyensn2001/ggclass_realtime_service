package notification

import (
	"context"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/enums"
	"ggclass_log_service/src/logger"
	"ggclass_log_service/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	CreateNotificationToUser(ctx context.Context, list []models.NotificationToUser) error
	GetNotifyByUserId(ctx context.Context, userId int) ([]models.NotificationToUser, error)
	FindByNotificationIds(ctx context.Context, ids []string) ([]models.Notification, error)
	FindByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.Notification, error)
	SetSeenForUser(ctx context.Context, userId int, notificationId string) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input createNotificationInput, typeNotification enums.NotificationType) (string, error) {
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
		Type:        typeNotification,
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

	pusher := config.GetConfig().Pusher

	err := pusher.Trigger("notifications", "refetch", users)
	if err != nil {
		logger.Sugar().Error(err)
	}

	logger.Sugar().Info(list)
	return s.repository.CreateNotificationToUser(ctx, list)

}

func (s *service) GetByUserId(ctx context.Context, userId int) ([]models.Notification, error) {

	logger.Sugar().Info(userId)

	notifyList, err := s.repository.GetNotifyByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	mapSeen := make(map[string]int)
	for _, item := range notifyList {
		ids = append(ids, item.NotificationId)
		mapSeen[item.NotificationId] = item.Seen
	}

	result, err := s.repository.FindByNotificationIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	for index, _ := range result {
		val, ok := mapSeen[result[index].ID.Hex()]
		if !ok {
			continue
		}
		result[index].Seen = val
	}

	return result, nil
}

func (s *service) GetByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.Notification, error) {
	return s.repository.FindByClassIdAndType(ctx, classId, typeNotification)
}

func (s *service) SetSeen(ctx context.Context, userId int, notificationId string) error {
	return s.repository.SetSeenForUser(ctx, userId, notificationId)
}
