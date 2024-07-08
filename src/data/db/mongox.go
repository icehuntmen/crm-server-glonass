package mongox

import (
	"context"
	"crm-glonass/src/config"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	MongoServer *mongo.Client
	DB          *mongo.Database
	ctx         context.Context
)

func Connection(conf *config.Config) {

	mongoUrl := fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s?authSource=%s`,
		conf.MongoX.Username, conf.MongoX.Password, conf.MongoX.Host, conf.MongoX.Port,
		conf.MongoX.Database, conf.MongoX.AuthSource)

	ctx = context.Background()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(mongoUrl)
	MongoServer, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		panic(err)
	}

	if err := MongoServer.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	DB := MongoServer.Database(conf.MongoX.Database)
	_, err = DB.Collection("settings").InsertOne(ctx, bson.D{{"key", "value"}})
	if err != nil {
		log.Fatal(err)
	}

	defer func(MongoServer *mongo.Client, ctx context.Context) {
		err := MongoServer.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(MongoServer, ctx)
}
