package bootstrap

import (
	"context"
	"fmt"
	"send2kobo/logger"
	"send2kobo/mongo"
	"time"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		logger.Error(err.Error())
	}

	err = client.Connect(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
	err = client.Ping(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Connection to MongoDB closed.")
}
