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

func (r *Redis) StoreSignal(ctx context.Context, vin string, time_stamp int64) error {
	return r.client.ZAdd(ctx, vin, redis.Z{
		Score:  float64(time_stamp),
		Member: fmt.Sprintf("%d", time_stamp),
	}).Err()
}

func (r *Redis) GetSignal(ctx context.Context, vin string) ([]string, error) {
	return r.client.ZRange(ctx, vin, 0, -1).Result()
}
