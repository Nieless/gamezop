package gamezop

import "gopkg.in/mgo.v2/bson"

type Game struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

// GameStore defines the operations of game store.
type GameStore interface {
	AddGame(game *Game) error
	GetGames() ([]*Game, error)
}
