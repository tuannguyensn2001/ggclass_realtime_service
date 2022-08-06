package notification

import (
	"context"
	notificationpb "ggclass_log_service/src/pb/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	Create(ctx context.Context, input createNotificationInput) (string, error)
	NotifyToUser(ctx context.Context, notificationId string, users []int) error
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

	id, err := t.service.Create(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, "has error")
	}

	return &notificationpb.CreateNotificationResponse{
		Message: "done",
		Id:      id,
	}, nil

}

func (t *transport) GetByUserId(ctx context.Context, request *notificationpb.GetNotificationByUserIdRequest) (*notificationpb.GetNotificationByUserIdResponse, error) {
	return &notificationpb.GetNotificationByUserIdResponse{
		Data: nil,
	}, nil
}
