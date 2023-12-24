package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBClient() *mongo.Client {
	conf := config.GetConfig()
	client, err := mongo.Connect(context.TODO(), buildConnectionOptions(conf))

	if err != nil {
		logger.Fatal("database", "NewMongoCollection", err.Error())
	}

	return client
}

func NewMongoCollection(client *mongo.Client) *mongo.Database {
	conf := config.GetConfig()
	c := client.Database(conf.DatabaseName)

	if c == nil {
		logger.Fatal("database", "NewMongoClient", "unable to retrieve collection")
	}

	return c
}

func buildConnectionOptions(conf config.Configuration) *options.ClientOptions {

	opt := options.Client()
	opt.SetConnectTimeout(time.Second * time.Duration(conf.MongoConnectionTimeout))
	opt.SetMaxPoolSize(uint64(conf.MongoMaxConnections))
	opt.SetMaxConnIdleTime(time.Second * 2)
	opt.SetReadPreference(readpref.PrimaryPreferred())

	if conf.MongoUser != "" && conf.MongoPassword != "" {
		opt.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    "admin",
			Username:      conf.MongoUser,
			Password:      conf.MongoPassword,
		})

		opt.SetHosts(strings.Split(conf.MongoHost, ","))

		return opt
	}
	return opt.ApplyURI(fmt.Sprintf("mongodb://%s", conf.MongoHost))
}
