package vehicle

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"brake_signal/controller"
	"brake_signal/service"
	"brake_signal/signal"
)

type Handlers struct {
	vehicleService *service.VehicleService
}

func NewHandlers(vehicleService *service.VehicleService) *Handlers {
	return &Handlers{
		vehicleService: vehicleService,
	}
}

func (h *Handlers) GenerateAddressPairs(c *gin.Context) {
	vin := c.Param("vin")
	if vin == "" {
		log.Error("handlers::GenerateAddressPairs, params error: nil vin")
		c.JSON(http.StatusBadRequest, controller.NewResult(nil, errors.New("empty vin")))
		return
	}
	vehicle, err := h.vehicleService.GetVehicle(c, vin)
	if err != nil {
		log.Errorf("handlers::GenerateAddressPairs, GetVehicle error:[%v], vin:[%v]", err, vin)
		c.JSON(http.StatusInternalServerError, controller.NewResult(nil, err))
		return
	}
	if vehicle != nil && vehicle.PublicKey != "" {
		log.Infof("handlers::GenerateAddressPairs, already create, vin:[%v], Pub:[%v]", vehicle.Vin, vehicle.PublicKey)
		c.JSON(http.StatusOK, controller.NewResult(vehicle, nil))
		return
	}
	v, err := h.vehicleService.CreateVehicle(c, vin)
	if err != nil {
		log.Errorf("handlers::GenerateAddressPairs, CreateVehicle error:[%v], vin:[%v]", err, vehicle.Vin)
		c.JSON(http.StatusInternalServerError, controller.NewResult(nil, err))
		return
	}
	log.Infof("handlers::GenerateAddressPairs, generate success, vin:[%v], Pub:[%v]", vehicle.Vin, vehicle.PublicKey)
	c.JSON(http.StatusOK, controller.NewResult(v, nil))
}

func (h *Handlers) VehicleSignalOnChain(c *gin.Context) {
	var req struct {
		Vin           string `json:"vin"`
		TimestampNano string `json:"time_stamp_nano"`
	}

	if err := c.BindJSON(&req); err != nil {
		log.Errorf("handlers::VehicleSignalToChain, GetVehicle error:[%v]", err)
		c.JSON(http.StatusInternalServerError, controller.NewResult(nil, err))
		return
	}

	if req.Vin == "" {
		log.Error("handlers::VehicleSignalToChain, vin invalid")
		c.JSON(http.StatusInternalServerError, controller.NewErrorResult(controller.INVALIDVIN))
		return
	}

	time_stamp_nano, err := strconv.ParseInt(req.TimestampNano, 10, 64)
	if err != nil {
		log.Errorf("handlers::VehicleSignalToChain, ParseInt error:[%v], str:[%v]", err, req.TimestampNano)
		c.JSON(http.StatusInternalServerError, controller.NewResult(nil, err))
		return
	}

	var resp struct {
		SignalInfo string `json:"signal_info"`
		SignalHash string `json:"signal_hash"`
		Status     string `json:"status"`
	}

	resp.SignalInfo, resp.SignalHash, err = h.vehicleService.OnChain(c, req.Vin, time_stamp_nano)
	if err != nil {
		log.Errorf("handlers::VehicleSignalToChain, OnChain error:[%v], vin:[%v], signal:[%v]", err, req.Vin, req.TimestampNano)
		c.JSON(http.StatusInternalServerError, controller.NewResult(nil, err))
		return
	}
	resp.Status = "上链中"

	c.JSON(http.StatusOK, controller.NewResult(resp, nil))
}

func (h *Handlers) GetVehicleSignalChainInfo(c *gin.Context) {
	vin := c.Param("vin")
	if vin == "" {
		log.Error("handlers::GetVehicleSignalChainInfo, params error: nil vin")
		c.JSON(http.StatusBadRequest, controller.NewResult(nil, errors.New("empty vin")))
		return
	}
	signal := signal.Signal{
		Vin:              "vin",
		Address:          "fdsfdsfewfdsfewrfw",
		LastOnChainBlock: 2023,
		LastOnChainHash:  "fdsfdsfdsfsdfs",
		LastOnChainInfo:  "1234567,1232434534,22343",
		CreateAt:         "1434645654",
		UpdateAt:         "1456546566",
	}
	c.JSON(http.StatusOK, controller.NewResult(signal, nil))
}
