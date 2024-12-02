package mongodb

import (
	"context"

	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/pkg/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context, cfg *config.Config) (*mongo.Collection, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)

	logger.Debug("create mongo client", constants.MongoCategory)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Error(err.Error(), constants.MongoCategory)
		return nil, err
	}

	db := client.Database(cfg.MongoDB)

	logger.DebugF("pinging mongo database (%s)", constants.MongoCategory, cfg.MongoDB)

	if err := db.RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		logger.Error(err.Error(), constants.MongoCategory)
		return nil, err
	}

	coll := db.Collection(cfg.MongoCollection)

	if _, err = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"expiration": 1},
		Options: options.Index().SetExpireAfterSeconds(cfg.MongoTTL),
	}); err != nil {
		logger.Error(err.Error(), constants.MongoCategory)
		return nil, err
	}

	logger.Info("successfully connected to MongoDB!!!", constants.MongoCategory)

	return coll, nil
}
