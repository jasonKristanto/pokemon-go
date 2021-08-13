package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations `mapstructure:"postgres"`
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Host string
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	Database    string `mapstructure:"database"`
	Host        string `mapstructure:"host"`
	Port    	int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	MaxConn     int    `mapstructure:"max_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	Timeout     int    `mapstructure:"timeout"`
}

var Configuration Configurations

func Init()  {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	PgSqlConnection()
}
