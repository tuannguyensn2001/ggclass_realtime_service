package notification

import (
	"context"
	"encoding/json"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/logger"
)

type rabbitTransport struct {
	service IService
}

func NewRabbitTransport(service IService) *rabbitTransport {
	return &rabbitTransport{service: service}
}

func (s *rabbitTransport) Bootstrap() {
	conn := config.GetConfig().RabbitMQ

	ch, err := conn.Channel()
	if err != nil {
		logger.Sugar().Error(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("notification", "direct", true, false, false, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}
	q1, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	err = ch.QueueBind(q.Name, "teacher_create", "notification", false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}
	err = ch.QueueBind(q1.Name, "seen", "notification", false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	msgs1, err := ch.Consume(q1.Name, "", true, false, false, false, nil)

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			content := d.Body

			var input notifyToUser
			err := json.Unmarshal(content, &input)
			if err == nil {
				err := s.service.NotifyToUser(context.Background(), input.NotificationId, input.Users)
				if err != nil {
					logger.Sugar().Error(err)
				}
			}
		}
	}()

	go func() {
		for d := range msgs1 {
			content := d.Body
			var input setSeenInput
			err := json.Unmarshal(content, &input)
			if err != nil {
				continue
			}

			err = s.service.SetSeen(context.Background(), input.UserId, input.NotificationId)
			if err != nil {
				logger.Sugar().Error(err)
			}

		}
	}()

	<-forever
}
