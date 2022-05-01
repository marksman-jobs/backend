package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDatabase(configuration Config) (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	option := options.Client().
		ApplyURI(configuration.MONGO_URI).
		SetMinPoolSize(uint64(configuration.MONGO_POOL_MIN)).
		SetMaxPoolSize(uint64(configuration.MONGO_POOL_MAX)).
		SetMaxConnIdleTime(time.Duration(configuration.MONGO_MAX_IDLE_TIME_SECOND) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	database := client.Database(configuration.MONGO_DATABASE)
	return database, nil

}
