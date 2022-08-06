package notification

import (
	"context"
	"ggclass_log_service/src/logger"
	"ggclass_log_service/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, notification *models.Notification) error {
	_, err := r.db.Database("ggclass_realtime").Collection(models.Notification{}.CollectionName()).InsertOne(ctx, notification)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateNotificationToUser(ctx context.Context, list []models.NotificationToUser) error {
	var documents []interface{}

	for _, item := range list {
		documents = append(documents, item)
	}

	_, err := r.db.Database("ggclass_realtime").Collection(models.NotificationToUser{}.CollectionName()).InsertMany(ctx, documents)
	if err != nil {
		logger.Sugar().Error(err)
		return err
	}

	return nil
}
