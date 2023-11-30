package vehicle

import (
	"brake_signal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.RouterGroup, vehicleService *service.VehicleService) {
	handlers := NewHandlers(vehicleService)
	r.GET("/pairs/:vin", handlers.GenerateAddressPairs)
}
