package common

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

const KeyringService = "dev.pixcee.strawhouse.command"
const KeyringUser = "strawhouse-command"

func InitConfig() {
	path := filepath.Join(os.Getenv("HOME"), ".config", "strawhouse")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal("Failed to create config directory: ", err)
	}
	viper.SetConfigName("command")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
}
