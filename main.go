package main

import (

	"github.com/grecinto/carinventory.git/apihandler"
	vehicleDA "github.com/grecinto/carinventory.git/vehicle_data_access"
	log "github.com/grecinto/carinventory.git/logger"
)

// logHandler is a wrapper function for logging, defaults to Logrus logging
var logHandler log.LoggerFunc = log.LogrusLog
// use the mock implementation for demo, 'will implement real DB when time permits.
var vehicleDataAccess vehicleDA.VehicleDataAccess = vehicleDA.NewMockAccess()

func main() {
	// initialize/read from .env/OS Env't the microservice(apihandler) config.
	config, err := apihandler.InitializeConfig()
	apih, err := apihandler.New("", *config, logHandler)
	if err != nil {
		// E.g. setup of HTTP ping endpoint (required by Kubernetes/AWS for health check) failed, exit right away.
		logHandler(err.Error(), log.Error)
		return
	}

	// Add the Vehicle Inventory API CRUD endpoint handlers
	apih.AddEndpoint("set", apihandler.POST, setVehicleHandler)
	apih.AddEndpoint("get", apihandler.GET, getVehicleHandler)

	// Support HTTP Verb DELETE for delete action, if time permits. :)
	//apih.AddEndpoint("delete", apihandler.DELETE, deleteVehicleHandler)

	// support delete on GET HTTP Verb for convenience. :) (easy to test using browser or postman)
	apih.AddEndpoint("delete", apihandler.GET, deleteVehicleHandler)

	// Listen & Serve requests
	apih.ListenAndServe()
}
