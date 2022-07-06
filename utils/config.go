package utils

import (
	"github.com/spf13/viper"
)

// type Dbconfig struct {
// 	Address  string `mapstructure:"address"`
// 	Port     string `mapstructure:"port"`
// 	Password string `mapstructure:"password"`
// 	User     string `mapstructure:"user"`
// }

// type ServerConfig struct {
// 	Port string `mapstructure:"port"`
// }

// type Config struct {
// 	Db     Dbconfig     `mapstructure:"db"`
// 	Server ServerConfig `mapstructure:"server"`
// }

type Dbconfig struct {
	Address string `map:"address"`
	Port    string `map:"port"`
}
type Config struct {
	DB Dbconfig `map:"db"`
}

func LoadConfig() (Config, error) {
	vp := viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./utils")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
