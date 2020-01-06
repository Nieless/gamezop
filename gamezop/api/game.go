package api

import (
	"github.com/gamezop/gamezop"
	"net/http"
)

type GameService struct {
	gamezop.GameStore
}

func (svc *GameService) GetGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		games, err := svc.GameStore.GetGames()
		if err != nil {
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}
		NewJSONWriter(w).Write(games, http.StatusOK)
	}
}
