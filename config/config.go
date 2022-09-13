package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() Configurations {

	fmt.Printf("Initializing config file \n")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Println("Port is\t\t", configuration.Server.Port)
	return configuration

}

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
	MODE         string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	MongoUser     string
	MongoPassword string
}
