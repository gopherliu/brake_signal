package service

import (
	"context"

	"brake_signal/vehicle"
)

type VehicleService struct {
	vehicleDB vehicle.DB
}

func NewVehicleService(vDB vehicle.DB) *VehicleService {
	return &VehicleService{
		vehicleDB: vDB,
	}
}

func (s *VehicleService) CreateVehicle(ctx context.Context, vin string) error {
	v := new(vehicle.Vehicle)
	v.Vin = vin
	return s.vehicleDB.Create(ctx, v)
}

func (s *VehicleService) GetVehicle(ctx context.Context, vin string) (*vehicle.Vehicle, error) {
	return s.vehicleDB.Get(ctx, vin)
}
