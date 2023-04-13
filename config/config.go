package config

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
	"log"
	"outgrow/dto"

	"github.com/spf13/viper"
)

type Config struct {
	Env string `mapstructure:"ENV"`

	Port   uint16 `mapstructure:"PORT"`
	Secret string `mapstructure:"SECRET" validate:"required,min=32"`

	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	PrivateKey ed25519.PrivateKey `mapstructure:"-"`
}

func GetConfig() *Config {
	var config Config

	viper.AddConfigPath("../../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("unable to load config file: %v", err)
		return nil
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = dto.Validate(config)
	if err != nil {
		log.Printf("failed to validate config: %v", err)
		return nil
	}

	err = config.GeneratePrivateKey()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &config
}

func (c *Config) GeneratePrivateKey() error {
	_, pk, err := ed25519.GenerateKey(bytes.NewReader([]byte(c.Secret)))
	if err != nil {
		return fmt.Errorf("secret unuseable")
	}

	c.PrivateKey = pk
	return nil
}
