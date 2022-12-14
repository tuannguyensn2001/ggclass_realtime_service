package notification

import (
	"context"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/enums"
	"ggclass_log_service/src/logger"
	"ggclass_log_service/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, notification *models.Notification) error {
	_, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.Notification{}.CollectionName()).InsertOne(ctx, notification)
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

	_, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.NotificationToUser{}.CollectionName()).InsertMany(ctx, documents)
	if err != nil {
		logger.Sugar().Error(err)
		return err
	}

	return nil
}

func (r *repository) GetNotifyByUserId(ctx context.Context, userId int) ([]models.NotificationToUser, error) {
	var result []models.NotificationToUser

	filter := bson.D{{"userId", userId}}

	cursor, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.NotificationToUser{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		logger.Sugar().Error(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var item models.NotificationToUser
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindByNotificationIds(ctx context.Context, ids []string) ([]models.Notification, error) {
	var result []models.Notification

	//filter := bson.D{{"_id", bson.D{{"$in", ids}}}}

	objectIds := make([]primitive.ObjectID, len(ids))

	for index, item := range ids {
		check, _ := primitive.ObjectIDFromHex(item)
		objectIds[index] = check
	}

	filter := bson.D{{"_id", bson.D{{"$in", objectIds}}}}

	opts := options.Find().SetSort(bson.D{{"_id", -1}})

	cursor, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.Notification{}.CollectionName()).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item models.Notification
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.Notification, error) {
	var result []models.Notification

	filter := bson.D{{"classId", classId}, {"type", typeNotification}}

	cursor, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.Notification{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item models.Notification
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) SetSeenForUser(ctx context.Context, userId int, notificationId string) error {

	filter := bson.D{{"notificationId", notificationId}, {"userId", userId}}
	update := bson.D{{"$set", bson.D{{"seen", 1}}}}

	_, err := r.db.Database(config.GetConfig().MongoDatabase).Collection(models.NotificationToUser{}.CollectionName()).UpdateOne(ctx, filter, update)

	return err
}
