package apihandler

import (
	"os"
	"strconv"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

//Config should contain the bare minimum values required by an api deployment environment.
// This is for use for by the Api Handler as it manages setting up the web endpoints.
type Config struct {
    Port                                   string `validate:"required,startswith=:"`
    LogLevel                               string `validate:"oneof=panic fatal error warn info debug trace"`
    ReadTimeoutSecs                        int 
    WriteTimeoutSecs                       int
    IdleTimeoutSecs						   int
}

// NewConfig instantiates Config with default settings.
func NewConfig() *Config{
	return &Config{
		Port: ":8080",
		LogLevel:"info",
		ReadTimeoutSecs: 300,
		WriteTimeoutSecs: 300,
		IdleTimeoutSecs:600,
	}
}

// Validate whether this config is valid or not, per annotations.
func (config Config) Validate() error {
	validate := validator.New()
	return validate.Struct(config)
}

// Getenv will return value of a given key, as read from "env".
func (c Config)Getenv(key string) string{
	return os.Getenv(key)
}

var envInitialized bool

// InitializeConfig will load OS "env" then issue a ReadConfig
// to actually read config from env and return config object.
func InitializeConfig() (*Config,error){
	if !envInitialized{
		envInitialized = true
		godotenv.Load()
	}
	return ReadConfig()
}
//ReadConfig provides a validated reading or loading of "bare minimum config"
//required by Api Handler. This reader relies on reading the "env" that is populated
//from the OS env or from the .env file for any key that
//doesn't already exist in the OS env. This allows a
//local .env to be a "fallback" for development.
func ReadConfig() (*Config,error){
	config := Config{}

	config.LogLevel = os.Getenv("LOG_LEVEL")
	config.Port = os.Getenv("PORT")
	config.ReadTimeoutSecs,_ = strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	config.WriteTimeoutSecs,_ = strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	config.IdleTimeoutSecs,_ = strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))

	// apply timeout defaults if too small read from config.
	if config.ReadTimeoutSecs < 5{
		config.ReadTimeoutSecs = 30
	}
	if config.WriteTimeoutSecs < 5{
		config.WriteTimeoutSecs = 30
	}
	if config.IdleTimeoutSecs < 10{
		config.IdleTimeoutSecs = 60
	}

	return &config, config.Validate()
}
