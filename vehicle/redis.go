package vehicle

import (
	"context"
	"fmt"

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

func (r *Redis) StoreSignal(ctx context.Context, key string, value int64) error {
	return r.client.ZAdd(ctx, key, redis.Z{
		Score:  float64(value),
		Member: fmt.Sprintf("%d", value),
	}).Err()
}

func (r *Redis) GetSignal(ctx context.Context, key string) ([]string, error) {
	return r.client.ZRange(ctx, key, 0, -1).Result()
}
