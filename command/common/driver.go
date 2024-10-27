package common

import (
	"github.com/bsthun/gut"
	"github.com/spf13/viper"
	"github.com/strawstacks/strawhouse-go"
	"github.com/zalando/go-keyring"
)

var Driver *strawhouse.Driver

func InitDriver() {
	key, err := keyring.Get(KeyringService, KeyringUser)
	if err != nil {
		gut.Fatal("failed to get key from keyring", err)
	}
	server := viper.Get("server")
	secure := viper.Get("secure")
	if server == nil {
		gut.Fatal("server is required, please use 'strawc config set --name server'", nil)
	}
	if secure == nil {
		gut.Fatal("secure is required, please use 'strawc config set --name secure' with value of <y/n>'", nil)
	}
	Driver, err = strawhouse.New(&strawhouse.Option{
		Key:    key,
		Server: server.(string),
		Secure: secure.(string) == "y",
	})
	if err != nil {
		gut.Fatal("failed to initialize driver", err)
	}
}
