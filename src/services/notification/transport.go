package notification

import (
	"context"
	"ggclass_log_service/src/enums"
	"ggclass_log_service/src/models"
	notificationpb "ggclass_log_service/src/pb/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IService interface {
	Create(ctx context.Context, input createNotificationInput, typeNotification enums.NotificationType) (string, error)
	NotifyToUser(ctx context.Context, notificationId string, users []int) error
	GetByUserId(ctx context.Context, userId int) ([]models.Notification, error)
	GetByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.Notification, error)
	SetSeen(ctx context.Context, userId int, notificationId string) error
}

type transport struct {
	notificationpb.UnimplementedNotificationServiceServer
	service IService
}

func NewTransport(service IService) *transport {
	return &transport{service: service}
}

func (t *transport) Create(ctx context.Context, request *notificationpb.CreateNotificationRequest) (*notificationpb.CreateNotificationResponse, error) {
	input := createNotificationInput{
		OwnerName:   request.OwnerName,
		OwnerAvatar: request.OwnerAvatar,
		CreatedBy:   int(request.CreatedBy),
		HtmlContent: request.HtmlContent,
		ClassId:     int(request.ClassId),
		Content:     request.Content,
	}

	id, err := t.service.Create(ctx, input, enums.NotificationType(request.Type))
	if err != nil {
		return nil, status.Error(codes.Internal, "has error")
	}

	return &notificationpb.CreateNotificationResponse{
		Message: "done",
		Id:      id,
	}, nil

}

func (t *transport) GetByUserId(ctx context.Context, request *notificationpb.GetNotificationByUserIdRequest) (*notificationpb.GetNotificationByUserIdResponse, error) {
	result, err := t.service.GetByUserId(ctx, int(request.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := make([]*notificationpb.Notification, len(result))

	for index, item := range result {
		data[index] = &notificationpb.Notification{
			OwnerAvatar: item.OwnerAvatar,
			OwnerName:   item.OwnerName,
			Id:          item.ID.Hex(),
			CreatedBy:   int64(item.CreatedBy),
			Content:     item.Content,
			HtmlContent: item.HtmlContent,
			ClassId:     int64(item.ClassId),
			CreatedAt:   timestamppb.New(*item.CreatedAt),
			UpdatedAt:   timestamppb.New(*item.UpdatedAt),
			Seen:        int64(item.Seen),
			Type:        int64(item.Type),
		}
	}

	return &notificationpb.GetNotificationByUserIdResponse{
		Data: data,
	}, nil
}

func (t *transport) GetByClassAndType(ctx context.Context, request *notificationpb.GetNotificationByClassAndTypeRequest) (*notificationpb.GetNotificationByClassAndTypeResponse, error) {
	result, err := t.service.GetByClassIdAndType(ctx, int(request.ClassId), enums.NotificationType(request.Type))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	data := make([]*notificationpb.Notification, len(result))

	for index, item := range result {
		data[index] = &notificationpb.Notification{
			OwnerAvatar: item.OwnerAvatar,
			OwnerName:   item.OwnerName,
			Id:          item.ID.Hex(),
			CreatedBy:   int64(item.CreatedBy),
			Content:     item.Content,
			HtmlContent: item.HtmlContent,
			ClassId:     int64(item.ClassId),
			CreatedAt:   timestamppb.New(*item.CreatedAt),
			UpdatedAt:   timestamppb.New(*item.UpdatedAt),
		}
	}

	return &notificationpb.GetNotificationByClassAndTypeResponse{Data: data}, nil

}
