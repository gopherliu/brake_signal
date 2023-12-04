package vehicle

import "context"

type Query struct {
}

type DB interface {
	Get(ctx context.Context, vin string) (*Vehicle, error)
	Create(ctx context.Context, v *Vehicle) error
	StoreSignal(ctx context.Context, vin string, time_stamp int64) error
	GetSignal(ctx context.Context, vin string) ([]string, error)
}
