package mongox

import (
	"context"
	"crm-glonass/config"
	"crm-glonass/pkg/logging"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoClient(conf *config.Config, ctx context.Context, logger logging.Logger) error {
	mongoUrl := fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s?authSource=%s`,
		conf.MongoX.Username, conf.MongoX.Password, conf.MongoX.Host, conf.MongoX.Port,
		conf.MongoX.Database, conf.MongoX.AuthSource)

	mongoconn := options.Client().ApplyURI(mongoUrl)
	var err error
	mongoClient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		logger.Fatal(logging.MongoDB, logging.Connection, err.Error(), nil)
	}
	return nil
}

func Execute(ctx context.Context, conf *config.Config, operation func(*mongo.Database) error) error {
	if mongoClient == nil {
		return errors.New("MongoDB client is not initialized")
	}

	db := mongoClient.Database(conf.MongoX.Database)
	return operation(db)
}

func Connection(conf *config.Config, ctx context.Context, logger logging.Logger) (*mongo.Database, error) {
	err := InitMongoClient(conf, ctx, logger)
	if err != nil {
		return nil, err
	}
	logger.Info(logging.MongoDB, logging.Connection, "Database connection established.", nil)
	return mongoClient.Database(conf.MongoX.Database), nil
}
