package main

import (
	"fmt"
	"github.com/gamezop/gamezop/mongo"
	"github.com/gamezop/gamezop/redis"
	gzSqs "github.com/gamezop/gamezop/sqs"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	mongoConnStr := MustGetEnv("MONGO_CONNECT_STRING")

	var session *mgo.Session
	session, err := mgo.Dial(mongoConnStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// redis client
	rClient, err := redis.NewClient(&redis.Config{Host:MustGetEnv("REDIS_HOST"), Port:MustGetEnv("REDIS_PORT")})
	if err != nil {
		panic(err)
	}

	sClient, err := gzSqs.NewClient(&gzSqs.Config{Region:MustGetEnv("AWSRegion")})
	if err != nil{
		panic(err)
	}

	qUrl := MustGetEnv("SQSQueueURL")
	sClient.QUrl = &qUrl

	gameDB := mongo.GameDB{DB:session.DB("")}

	qService := &QueueService{
		MongoGameStore:&gameDB,
		RedisGameStore:rClient,
		SQSStore:sClient,
	}

	for {
		qService.UpdateDBStores()
	}
}


// MustGetEnv gets an environment variable or panics.
func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s missing", key))
	}
	return v
}