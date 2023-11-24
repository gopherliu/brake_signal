package signal

import "github.com/redis/go-redis/v9"

type Redis struct {
	client *redis.Client
}

func NewRedis(c *redis.Client) *Redis {
	return &Redis{
		client: c,
	}
}
