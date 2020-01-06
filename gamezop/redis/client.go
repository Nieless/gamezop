package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Host string
	Port string
}

// Client is redis client
type Client struct {
	rClient *redis.Client
}

func NewClient(config *Config)(*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Host + ":" + config.Port,
		Password:     "", // no password set
		DB:           0,  // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return &Client{client}, nil
}