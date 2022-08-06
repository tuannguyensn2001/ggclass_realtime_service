package routes

import (
	"context"
	notificationpb "ggclass_log_service/src/pb/notification"
	"ggclass_log_service/src/services/notification"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(server *grpc.Server) {
	//assignmentpb.RegisterLogAssignmentServiceServer(server, buildAssignmentTransport())
	notificationpb.RegisterNotificationServiceServer(server, buildNotificationTransport())
}

func RegisterGrpcGateway(ctx context.Context, m *runtime.ServeMux, conn *grpc.ClientConn) error {
	//err := assignmentpb.RegisterLogAssignmentServiceHandler(ctx, m, conn)
	//if err != nil {
	//	return err
	//}
	err := notificationpb.RegisterNotificationServiceHandler(ctx, m, conn)
	if err != nil {
		return nil
	}

	return nil
}

//func buildAssignmentTransport() assignmentpb.LogAssignmentServiceServer {
//
//	transport := assignment.NewTransport(assignment.BuildService())
//
//	rabbitTransport := assignment.NewRabbitTransport(assignment.BuildService())
//	rabbitTransport.Bootstrap(context.Background())
//
//	return transport
//
//}

func buildNotificationTransport() notificationpb.NotificationServiceServer {
	transport := notification.NewTransport(notification.BuildService())
	rabbitTransport := notification.NewRabbitTransport(notification.BuildService())
	go rabbitTransport.Bootstrap()
	return transport
}
