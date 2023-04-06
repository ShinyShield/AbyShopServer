package config

import (
	"strings"

	"github.com/spf13/viper"
)

var c config

type config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Username     string
	Password     string
	Name         string
	Host         string
	Port         uint64
	Encoding     string
	Maxconns     uint64
	Maxidleconns uint64
	Timeout      uint64
	Logmode      bool
}

// init database
func Init() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigName("config")
	v.AddConfigPath(".")

	/*err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}*/

	v.Unmarshal(&c)
}

func GetDatabaseConfig() DatabaseConfig {
	return c.Database
}
