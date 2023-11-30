package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"brake_signal/config"
	vehicleController "brake_signal/controller/vehicle"
	"brake_signal/service"
	"brake_signal/vehicle"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	var confFile string
	flag.StringVar(&confFile, "c", "", "conf filename")
	flag.Parse()
	var err error
	var jsonBlob []byte
	if jsonBlob, err = os.ReadFile(confFile); err != nil {
		log.Fatalf("read file error:[err]", err)
	}
	if err = json.Unmarshal(jsonBlob, &config.C); err != nil {
		log.Fatalf("Unmarshal error:[err]", err)
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.C.RedisAddr,
		Password: config.C.RedisAuth, // no password set
		DB:       0,                  // use default DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	vDB := vehicle.NewRedis(client)
	v_svc := service.NewVehicleService(vDB)
	vehicleController.NewHandlers(v_svc)

	r := gin.Default()
	g := r.Group("/api/v1")
	vehicleController.NewRouter(g, v_svc)

	r.Run(config.C.Bind) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
