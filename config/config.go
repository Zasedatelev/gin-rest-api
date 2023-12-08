package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
	JWT      JWTConfig      `yaml:"jwt"`
}

type ServerConfig struct {
	InternalPort int `yaml:"internalPort"`
}

type PostgresConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	SSlMode  string `yaml:"sslMode"`
}

type JWTConfig struct {
	Secret                    string        `yaml:"secret"`
	AccessTokenExpireDuration time.Duration `yaml:"accessTokenExpireDuration"`
	// RefreshTokenExpireDuration time.Duration `yaml: " " `
	// RefreshSecret              string        `yaml: " " `
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			InternalPort: 8080,
		},
		Postgres: PostgresConfig{
			User:     "",
			Password: "",
			DbName:   "",
			SSlMode:  "",
		},
		JWT: JWTConfig{
			Secret:                    "",
			AccessTokenExpireDuration: 60,
		},
	}
}

func LoadConfig() *Config {
	config := NewConfig()
	yamlFile, err := os.ReadFile("./config/config-dev.yml")
	if err != nil {
		log.Fatalf("#ERROR: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("#ERROR: %v", err)
	}
	return config

}
