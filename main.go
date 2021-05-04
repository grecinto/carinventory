package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/grecinto/carinventory.git/apihandler"
	vehicleDA "github.com/grecinto/carinventory.git/vehicle_data_access"
	log "github.com/grecinto/carinventory.git/logger"
)

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

// logHandler is a wrapper function for logging, defaults to Logrus logging
var logHandler log.LoggerFunc = log.LogrusLog

// use the mock implementation for demo, 'will implement real DB when time permits.
var vehicleDataAccess vehicleDA.VehicleDataAccess = vehicleDA.NewMockAccess()

// setVehicleHandler creates or updates a set of Vehicles
// Returns:
// 200 (OK) - item(s) created/updated, 500 error or request has no content 
func setVehicleHandler(w http.ResponseWriter, r *http.Request) {
	var vehicles []vehicleDA.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicles)
	if err != nil{
		writeResponse(w, r, "Parameter decoding failure", apihandler.ServerError, err)
		return
	}
	validationResult := vehicleDA.Validate(vehicles...)
	if len(validationResult) > 0 {
		writeResponse(w, r, "Item(s) failed Validation", apihandler.ServerError)
		return
	}

	// all or nothing(perhaps we want to support partial on non-demo version).
	// only process if payload are all valid Vehicle structs

	// TODO: log result Errors when real backend gets implemented and time permits.

	results := vehicleDataAccess.Set(vehicles...)
	if len(results) == 0 {
		writeResponse(w, r, "Item(s) Upserted successfully", apihandler.ServerSuccess)
		return
	}
	if len(results) == len(vehicles) {
		writeResponse(w, r, "Item(s) Upsert failed, see server logs for details", apihandler.ServerError)
		return
	}
	writeResponse(w, r, "Item(s) Upsert partially succeeded", apihandler.ServerSuccess)
}

// getVehicleHandler retrieves and returns a set of Vehicles given their VIN numbers.
// Returns:
// 200 (OK) - item(s) retrieved from DB, 404 - no item was found, 500 error or request has no content
// Sample HTTP GET URL: http://localhost:8080/get?VIN=1234&VIN=1235
func getVehicleHandler(w http.ResponseWriter, r *http.Request) {
    vins,ok := r.URL.Query()["VIN"]
	if !ok || len(vins) < 1{
		writeResponse(w, r, "Parameter 'VIN' array can't be extracted from URL", apihandler.ServerError)
		return
	}
	results := vehicleDataAccess.Get(vins...)
	if vehicleDA.AllNotFound(results){
		writeResponse(w, r, "No Item was found", apihandler.ItemNotFound)
		return
	}
	data, err := json.Marshal(results)
	if err != nil {
		writeResponse(w, r, "Results marshal failed", apihandler.ServerError, err)
		return
	}

	// write retrieved Items' data to response stream.
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(apihandler.ServerSuccess)
	w.Write(data)
}

// deleteVehicleHandler will delete a set of Vehicles from the backend DB, given their VIN numbers.
// Returns:
// 200 (OK) - item(s) deleted in DB, 404 - no item was found for delete, 500 error or request has no content
// Sample URL: http://localhost:8080/delete?VIN=1234&VIN=1235
func deleteVehicleHandler(w http.ResponseWriter, r *http.Request) {
    vins,ok := r.URL.Query()["VIN"]
	if !ok || len(vins) < 1{
		writeResponse(w, r, "Parameter 'VIN' array can't be extracted from URL", apihandler.ServerError)
		return
	}
	results := vehicleDataAccess.Delete(vins...)
	if vehicleDA.DeleteAllFailed(results){
		writeResponse(w, r, "No Item with submitted VINs was found for delete", apihandler.ItemNotFound)
		return
	}

	data, err := json.Marshal(results)
	if err != nil {
		writeResponse(w, r, "Results marshal failed", apihandler.ServerError, err)
		return
	}
	if vehicleDA.DeleteAllSucceeded(results){
		writeResponse(w, r, "Item(s) with VINs submitted were all deleted", apihandler.ServerSuccess)
		return
	}

	// write delete results to response stream. (partial delete)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(apihandler.ServerSuccess)
	w.Write(data)
}

func writeResponse(w http.ResponseWriter, r *http.Request, data string, serverReturnCode int, errors ...error){
	w.WriteHeader(serverReturnCode)
	w.Write([]byte(data))

	// TODO: log some info about the request when time permits

	if len(errors) > 0 {
		logHandler(fmt.Sprintf("Server return code: %d", serverReturnCode), log.Error)
		for _,e := range errors{
			logHandler(e.Error(), log.Error)
		}
		return
	}
	logHandler(fmt.Sprintf("Server return code: %d", serverReturnCode), log.Debug)
}
