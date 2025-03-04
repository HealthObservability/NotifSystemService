package config

import (
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres   Postgres   `yaml:"postgres"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type Postgres struct {
	Host     string `yaml:"POSTGRES_HOST" validate:"required"`
	Port     string `yaml:"POSTGRES_PORT" validate:"required"`
	User     string `yaml:"POSTGRES_USER" validate:"required"`
	Password string `yaml:"POSTGRES_PASSWORD" validate:"required"`
	Database string `yaml:"POSTGRES_DB" validate:"required"`
}

type HTTPServer struct {
	Port         string   `yaml:"HTTP_PORT" validate:"required"`
	AllowOrigins []string `yaml:"ALLOW_ORIGINS"`
}

// MustConfigure is a configurator for a config.
// @params: configPath is a path to config. is used by tests from any package with configuration.
func MustConfigure(configPath string) Config {
	log := logger.Logger.WithField("op", "config.MustConfigure")

	config := &Config{}
	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		description, _ := cleanenv.GetDescription(config, nil)
		log.Fatalf("failed to read config %s %s", description, err)
	}

	validate := validator.New()

	if err = validate.Struct(*config); err != nil {
		log.Fatal(err)
	}

	return *config
}
