package main

import (
	"github.com/gamezop/gamezop"
	"github.com/gamezop/gamezop/redis"
	gamezopSqs "github.com/gamezop/gamezop/sqs"
	"log"
)

type QueueService struct {
	MongoGameStore gamezop.GameStore
	RedisGameStore redis.GameStore
	SQSStore       gamezopSqs.SqsStore
}

func (svc *QueueService) UpdateDBStores() {
	msgs, err := svc.SQSStore.ReceiveMessages()
	if err != nil || len(msgs) == 0{
		svc.UpdateDBStores()
	}

	for receiptHandle, msg := range msgs{

		// add game entry from redis
		game, err := svc.RedisGameStore.GetGameByKey(*msg)
		if err != nil{
			log.Println(err.Error())
			continue
		}

		// add game entry to mongo
		if err = svc.MongoGameStore.AddGame(game); err != nil{
			log.Println(err.Error())
			continue
		}

		// delete from redis
		svc.RedisGameStore.DeleteGameByKey(*msg)

		// delete from queue, ignore err
		svc.SQSStore.DeleteMsg(&receiptHandle)
	}
}
