package main

import (
	"fmt"
	"github.com/gamezop/gamezop/api"
	"github.com/gamezop/gamezop/mongo"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
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

	gameAPI := &api.GameService{
		GameStore: &mongo.GameDB{DB: session.DB("")},
	}

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/games", gameAPI.GetGames())

	err = ListenAndServe(MustGetEnv("HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}

// ListenAndServe serves the application.
func ListenAndServe(port string, handler http.Handler) error {
	fmt.Println("Listening on:", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

// MustGetEnv gets an environment variable or panics.
func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s missing", key))
	}
	return v
}
