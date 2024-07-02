package service

import (
	common "github.com/LerianStudio/midaz/common"
)

// Config is the top level configuration struct for the entire application.
type Config struct {
	EnvName           string `env:"ENV_NAME"`
	ServerAddress     string `env:"SERVER_ADDRESS"`
	PrimaryDBHost     string `env:"DB_HOST"`
	PrimaryDBUser     string `env:"DB_USER"`
	PrimaryDBPassword string `env:"DB_PASSWORD"`
	PrimaryDBName     string `env:"DB_NAME"`
	PrimaryDBPort     string `env:"DB_PORT"`
	ReplicaDBHost     string `env:"DB_REPLICA_HOST"`
	ReplicaDBUser     string `env:"DB_REPLICA_USER"`
	ReplicaDBPassword string `env:"DB_REPLICA_PASSWORD"`
	ReplicaDBName     string `env:"DB_REPLICA_NAME"`
	ReplicaDBPort     string `env:"DB_REPLICA_PORT"`
	MongoDBHost       string `env:"MONGO_HOST"`
	MongoDBName       string `env:"MONGO_NAME"`
	MongoDBUser       string `env:"MONGO_USER"`
	MongoDBPassword   string `env:"MONGO_PASSWORD"`
	MongoDBPort       string `env:"MONGO_PORT"`
	AuthHost          string `env:"AUTH_HOST"`
	AuthPort          string `env:"AUTH_PORT"`
	AuthRealm         string `env:"AUTH_REALM"`
	AuthEndpoint      string `env:"AUTH_ENDPOINT"`
	AuthClientID      string `env:"AUTH_CLIENT_ID"`
	AuthClientSecret  string `env:"AUTH_CLIENT_SECRET"`
}

// NewConfig creates a instance of Config.
func NewConfig() *Config {
	cfg := &Config{}

	if err := common.SetConfigFromEnvVars(cfg); err != nil {
		panic(err)
	}

	return cfg
}
