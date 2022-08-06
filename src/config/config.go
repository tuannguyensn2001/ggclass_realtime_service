package config

import (
	"context"
	"ggclass_log_service/src/logger"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type structure struct {
	App struct {
		HttpPort string `mapstructure:"httpPort"`
		GrpcPort string `mapstructure:"grpcPort"`
	} `mapstructure:"app"`
	Mongo struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"mongo"`
	RabbitMQ struct {
		Url string `mapstructure: "url"`
	} `mapstructure:"rabbitmq"`
}

type config struct {
	HttpPort string
	GrpcPort string
	Mongo    *mongo.Client
	RabbitMQ *amqp091.Connection
}

var cfg config

func GetConfig() config {
	return cfg
}

func Load() error {
	path, _ := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config := structure{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Mongo.Url))
	if err != nil {
		logger.Sugar().Errorf("err connect mongo", err)
	}

	rabbit, err := amqp091.Dial(config.RabbitMQ.Url)
	if err != nil {
		logger.Sugar().Error(err)
	}

	cfg.HttpPort = config.App.HttpPort
	cfg.GrpcPort = config.App.GrpcPort
	cfg.Mongo = client
	cfg.RabbitMQ = rabbit

	return nil

}
