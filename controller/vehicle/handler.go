package vehicle

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"brake_signal/controller"
	"brake_signal/service"
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
