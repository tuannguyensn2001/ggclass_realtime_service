package notification

import "ggclass_log_service/src/config"

func BuildService() *service {
	repository := NewRepository(config.GetConfig().Mongo)
	service := NewService(repository)
	return service
}
