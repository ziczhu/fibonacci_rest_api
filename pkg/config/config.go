package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var config *Config
var once sync.Once

// Config for Application
type Config struct {
	Env string `envconfig:"ENV" default:"dev"`

	Port int `envconfig:"PORT" default:"8080"`

	// MaxFibInput defines the max fibonacci number will be allowed to calculate
	MaxFibInput int `envconfig:"MAX_FIB_INPUT" default:"10000"`
	// InitFibCacheSize defines the initial size the cache
	InitFibCacheSize int `envconfig:"INIT_FIB_CACHE_SIZE" default:"1000"`
	// MaxFibCacheSize defines the max size the cache
	MaxFibCacheSize int `envconfig:"MAX_FIB_CACHE_SIZE" default:"5000"`
}

func InitConfig() *Config {
	once.Do(func() {
		var c Config
		err := envconfig.Process("", &c)
		if err != nil {
			panic(err)
		}
		config = &c
	})
	return config
}

// GetConfig return the singleton
func GetConfig() *Config {
	if config == nil {
		config = InitConfig()
	}
	return config
}
