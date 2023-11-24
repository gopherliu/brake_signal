package vehicle

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(c *redis.Client) *Redis {
	return &Redis{
		client: c,
	}
}

func (r *Redis) Create(ctx context.Context, v *Vehicle) error {
	return r.client.HMSet(ctx, v.Vin, v).Err()
}

func (r *Redis) Get(ctx context.Context, vin string) (*Vehicle, error) {
	v := new(Vehicle)
	err := r.client.HGetAll(ctx, vin).Scan(v)
	return v, err
}
