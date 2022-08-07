package config

import (
	"context"
	"ggclass_log_service/src/logger"
	"github.com/pusher/pusher-http-go/v5"
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
	Pusher   pusher.Client
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

	pusherClient := pusher.Client{
		AppID:   "1440558",
		Key:     "26bdb6fd186156c41fe6",
		Secret:  "708f13f675065ba00a92",
		Cluster: "ap1",
		Secure:  true,
	}

	cfg.HttpPort = config.App.HttpPort
	cfg.GrpcPort = config.App.GrpcPort
	cfg.Mongo = client
	cfg.RabbitMQ = rabbit
	cfg.Pusher = pusherClient

	return nil

}
