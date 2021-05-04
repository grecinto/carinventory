package apihandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/grecinto/carinventory.git/logger"
)

// Handler
type Handler struct {
	Router *mux.Router
	Config Config
	loggerFunc logger.LoggerFunc
	pathPrefix string
}

// ServerReturnCodes
const (
	ServerSuccess = 200
	ServerError = 400
	ItemNotFound = 404
)

// HTTPVerb string enumeration, e.g. POST, GET, DELETE
type HTTPVerb string
const(
	POST HTTPVerb = "POST"
	GET HTTPVerb = "GET"
	DELETE HTTPVerb = "DELETE"
)

// New returns the package api handler containing members necessary for
// setting up web endpoint(s). This expects "env" to contain minimum required configurations of a Handler.
// See Config.go for list of these fields.
func New(pathPrefix string, config Config, logHandler logger.LoggerFunc) (*Handler,error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// set global log level to the level set in config.
	logger.SetLogLevel(config.LogLevel)

	handler := &Handler{
		Config: config,
		loggerFunc: logHandler,
		pathPrefix: pathPrefix,
	}
	handler.Router = handler.newRouter()

	handler.log("API Host Web Service started.", logger.Info)

	return handler, nil
}

// AddEndpoint will add an endpoint request handler for a given action Verb(POST, GET pr DELETE)
func (handler *Handler) AddEndpoint(name string, httpVerb HTTPVerb, requestHandler func (w http.ResponseWriter, r *http.Request)) {
	if handler.pathPrefix != ""{
		handler.Router.PathPrefix(handler.pathPrefix).Subrouter().HandleFunc("/" + name, requestHandler).Methods(string(httpVerb))
		return
	}
	handler.Router.HandleFunc("/" + name, requestHandler).Methods(string(httpVerb))
}

// ServeHTTP will route a request to an endpoint.
func (handler *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	handler.Router.ServeHTTP(w,req)
}

// ListenAndServe launch the web service and start listening & delegation of requests to endpoint handlers.
func (handler *Handler) ListenAndServe() error{
	if handler.Router == nil{
		return fmt.Errorf("mux.Router is nil, a typical problem is failure reading the base config values(env entries missing?).")
	}
	handler.log("Started to listen and serve on Port " + handler.Config.Port, logger.Info)

	//Start Handler and Graceful Shutdown
	srv := &http.Server{
		Addr:    handler.Config.Port,
		Handler: handler.Router,
		WriteTimeout: time.Second * time.Duration(handler.Config.WriteTimeoutSecs),
		ReadTimeout:  time.Second * time.Duration(handler.Config.ReadTimeoutSecs),
		IdleTimeout:  time.Second * time.Duration(handler.Config.IdleTimeoutSecs),
	}
	err := listenAndServe(srv)
	return err
}

// initialize the Router and pre-create the default health check(ping) endpoint
func (h *Handler)newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", h.pingHandler).Methods(string(GET))

	// do we want to log each request call? no need for a demo/interview thing.
	//r.Use(log.LogHandler)

	return r
}

//Enables simple health check during deployment
func (h *Handler)pingHandler(w http.ResponseWriter, r *http.Request) {
	h.log("Received ping", logger.Info)
	w.Write([]byte("pong"))
}

// just for convenience, will delegate call to actual log function received in CTOR
func (h *Handler) log(logMessage string, logLevel logger.LogLevel){
	if h.loggerFunc != nil{
		h.loggerFunc(logMessage, logLevel)
	}
}
