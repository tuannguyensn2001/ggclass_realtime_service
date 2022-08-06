package main

import (
	"context"
	"ggclass_log_service/src/cmd"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/logger"
	"log"
)

func main() {

	config.Load()

	logger.InitLog()
	defer func() {
		logger.SyncLog()
		config.GetConfig().Mongo.Disconnect(context.Background())
		config.GetConfig().RabbitMQ.Close()
	}()

	rootCmd := cmd.GetRoot()

	err := rootCmd.Execute()

	if err != nil {
		log.Fatalln(err)
	}
}
