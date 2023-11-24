package vehicle

import "context"

type Query struct {
}

type DB interface {
	Get(ctx context.Context, vin string) (*Vehicle, error)
	Create(ctx context.Context, v *Vehicle) error
}
