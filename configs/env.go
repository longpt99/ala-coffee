package configs

import (
	"ecommerce/common/models"
	"fmt"

	"github.com/spf13/viper"
)

var Env models.Config

func InitConfig() error {
	viper.SetConfigFile("dev.env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		panic(fmt.Errorf("fatal unmarshal config file: %w", err))
	}

	return nil
}
