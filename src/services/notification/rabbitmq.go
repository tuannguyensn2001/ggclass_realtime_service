package notification

import (
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

	err = ch.QueueBind(q.Name, "teacher_create", "notification", false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			logger.Sugar().Info(string(d.Body))
		}
	}()
	<-forever
}
