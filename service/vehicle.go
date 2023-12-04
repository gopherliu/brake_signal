package service

import (
	"context"
	"errors"
	"strings"
	"time"

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

func (s *VehicleService) OnChain(ctx context.Context, vin string, time_stamp int64) (string, string, error) {
	if time_stamp > time.Now().UnixMicro() {
		log.Errorf("VehicleService::OnChain, time error, vin:[%v], time_stamp:[%v], now:[%v]",
			vin, time_stamp, time.Now().UnixMicro())
		return "", "", errors.New("invalid time stamp nano")
	}

	err := s.vehicleDB.StoreSignal(ctx, vin, time_stamp)
	if err != nil {
		log.Errorf("VehicleService::OnChain, StoreSignal error:[%v], vin:[%v]", err, vin)
		return "", "", err
	}

	signals, err := s.vehicleDB.GetSignal(ctx, vin)
	if err != nil {
		log.Errorf("VehicleService::OnChain, GetSignal error:[%v], vin:[%v]", err, vin)
		return "", "", err
	}

	signal_string := strings.Join(signals, ",")
	log.Infof("VehicleService::OnChain, signal_string:[%v]", signal_string)
	signal_hash := utils.GenerateHash256(signal_string)
	log.Infof("VehicleService::OnChain, signal_hash:[%v]", signal_hash)

	// TODO::上链
	return signal_string, signal_hash, err
}
