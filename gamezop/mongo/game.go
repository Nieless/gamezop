package mongo

import (
	"github.com/gamezop/gamezop"
	"gopkg.in/mgo.v2"
)

type GameDB struct {
	DB *mgo.Database
}

// GetGames gets a Mongo games document from Mongo
func (mgo *GameDB) GetGames() ([]*gamezop.Game, error) {
	// Create variable for collection
	cGameConfig := mgo.DB.C("games")
	games := make([]*gamezop.Game, 0)
	if err := cGameConfig.Find(nil).All(&games); err != nil {
		return nil, err
	}

	return games, nil
}

// AddGame adds a game document to games collection in Mongo
func (mgo *GameDB) AddGame(game *gamezop.Game) error {
	// Create variable for collection
	cGameConfig := mgo.DB.C("games")
	return cGameConfig.Insert(&game)
}

// DeleteGame deletes a Game document from Mongo
func (mgo *GameDB) DeleteGame(game gamezop.Game) error {
	// Create variable for collection
	cGameConfig := mgo.DB.C("games")
	return cGameConfig.RemoveId(game.ID.Hex())
}
