package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"quick_gin/config/env"
	"time"
)

var mongoClient *mongo.Client
var mongoIsConnection = false

func connectMongodb() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(env.MongoDBConfig.Timeout)*time.Second)
	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI("mongodb://"+env.MongoDBConfig.Host).
			SetAuth(options.Credential{
				Username: env.MongoDBConfig.Username,
				Password: env.MongoDBConfig.Password,
			}))
	if err != nil {
		log.Fatalln("MongoDB连接失败 " + err.Error())
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("MongoDB无法Ping通 " + err.Error())
	}
	mongoIsConnection = true
}

//记录日志到mongodb中
func Instance() *mongo.Client {
	if !mongoIsConnection {
		connectMongodb()
	}
	return mongoClient
}
