package utils

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)


func LoadEnv(configFileName string)  {
	// name of config file (without extension)
	//".env"
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

// return the value of the key
func EnvVar(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func EnvVarArray(key string) []string {
	value := viper.GetStringSlice(key)
	if len(value) == 0 {
		fmt.Errorf("environment variable '%s' not found", key)
		return nil
	}
	return value
}