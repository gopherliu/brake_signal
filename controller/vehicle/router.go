package vehicle

import (
	"net/http"

	"brake_signal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.RouterGroup, vehicleService *service.VehicleService) {
	handlers := NewHandlers(vehicleService)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pairs/:vin", handlers.GenerateAddressPairs)
	r.POST("/onChain", handlers.VehicleSignalOnChain)
	r.GET("/getOnChainInfo/:vin", handlers.GetVehicleSignalChainInfo)
}
