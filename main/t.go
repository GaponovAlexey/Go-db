package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	vp := viper.New()
	vp.SetConfigName("test")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vp.GetString("foo"))
}
