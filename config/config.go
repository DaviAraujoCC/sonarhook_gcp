package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	QUALITY_GATE_STATUS_FILTER = "quality_gate_status"
	GOOGLE_CHAT_WEBHOOK_URL = "google_chat_webhook_url"
)

var (
	// Timezone is the timezone of the server
	Timezone string
	// Config file path
	cfgFile = os.Getenv("CONFIG_FILE")
)

type Config struct {
	Timezone string `mapstructure:"timezone"`
	Webhooks []Webhook 
}

type Webhook struct {
	Path string `mapstructure:"path"`
	Parameters map[string]string `mapstructure:"parameters"`
}

func NewConfig() *Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config.yaml")
	}

	viper.SetConfigType(filepath.Ext(cfgFile)[1:])

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		
	} else {
		log.Fatal(err)
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	switch {
	case viper.GetString("config.timezone") == "":
		log.Println("TIMEZONE is not set, using default America/Sao_Paulo")
		Timezone = "America/Sao_Paulo"
	}

	return cfg
}