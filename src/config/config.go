package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	SSLMode            bool
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config{
	cfgPath := getConfigpath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")

	if err != nil {
		log.Fatalf("error in load config: %v", err)

	}
	cfg, err := ParsConfig(v)

	if err != nil {
		log.Fatalf("error in parse config: %v", err)

	}

	return cfg

}

func ParsConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to Parse Config %s", err)
		return nil, err
	}
	return &cfg, err

} // tabdil be struc config

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil {
		log.Printf("unable to read Config %s", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")

		}
		return nil, err
	}
	return v, nil

} //gereft file ro va tabdil kard be struct viper

func getConfigpath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "production" {
		return "config/Config-production"

	} else {
		return "../config/config-development"

	}

} //gereftan file
