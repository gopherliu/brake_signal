package vehicle

import (
	"brake_signal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, vehicleService *service.VehicleService) *gin.RouterGroup {
	handlers := NewHandlers(vehicleService)
	r.POST("/api/v1/pairs/:vin", handlers.GenerateAddressPairs)
	g := r.Group("/api/v1")
	return g
}
