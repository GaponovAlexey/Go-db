package ut

import "github.com/spf13/viper"

// type SetConfigFile struct {
// 	Name string `map:"name"`
// }

// type Config struct {
// 	DB SetConfigFile `map:"db"`
// }
type Dbconfig struct {
	Address string `map:"address"`
	Port    string `map:"port"`
}
type Config struct {
	DB Dbconfig `map:"db"`
}

func GetConfigJson() (Config, error) {
	vp := viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./ut")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, err
}
