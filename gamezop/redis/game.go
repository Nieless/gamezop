package redis

import (
	"encoding/json"
	"github.com/gamezop/gamezop"
	"gopkg.in/mgo.v2/bson"
)

// GameStore defines the operations of game store.
type GameStore interface {
	GetGameByKey(key string) (*gamezop.Game, error)
	DeleteGameByKey(key string)
}

// GetGameByKey gets entry of game from redis
func (rc *Client) GetGameByKey(key string) (*gamezop.Game, error) {
	dataStr, err := rc.rClient.Get(key).Result()
	if err != nil {
		return nil, err
	}
	
	game := &gamezop.Game{}
	game.ID = bson.NewObjectId()

	err = json.Unmarshal([]byte(dataStr), game)
	if err != nil {
		return nil, err
	}

	return game, err
}

// GetGameByKey deletes entry of game from redis
func (rc *Client) DeleteGameByKey(key string) {
	// TODO handle err here
	_ = rc.rClient.Del(key)
	return
}
