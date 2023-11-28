package service

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"brake_signal/utils"
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

func (s *VehicleService) CreateVehicle(ctx context.Context, vin string) (*vehicle.Vehicle, error) {
	var (
		err error
	)
	v := new(vehicle.Vehicle)
	v.Vin = vin
	v.PrivateKey, v.PublicKey, err = utils.GenerateKeyPair()
	if err != nil {
		log.Errorf("VehicleService::CreateVehicle, GenerateKeyPair error:[%v], vin:[%v]", err, v.Vin)
		return nil, err
	}
	if v.PrivateKey == "" || v.PublicKey == "" {
		log.Errorf("VehicleService::CreateVehicle, pairs nil, vin:[%v]", v.Vin)
		return nil, errors.New("invalid pairs")
	}
	return v, s.vehicleDB.Create(ctx, v)
}

func (s *VehicleService) GetVehicle(ctx context.Context, vin string) (*vehicle.Vehicle, error) {
	return s.vehicleDB.Get(ctx, vin)
}
