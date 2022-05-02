package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Databases struct {
	Mongo    *mongo.Database
	Postgres *pgxpool.Pool
}

func InitializeDatabase(configuration Config) (*Databases, error) {

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

	mongo := client.Database(configuration.MONGO_DATABASE)

	pgconn, err := pgxpool.Connect(ctx, configuration.POSTGRES_URI)
	if err != nil {
		return nil, err
	}
	defer pgconn.Close()

	database := &Databases{
		Mongo:    mongo,
		Postgres: pgconn,
	}

	return database, nil

}
