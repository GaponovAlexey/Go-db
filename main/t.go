package main

import (
	"conect-db/utils"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("start")

	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)

	vp := viper.New()
	vp.SetConfigName("test")
	vp.SetConfigType("json")
	vp.AddConfigPath(".test")
	// err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vp.GetString("name"))

	vp.Set("Nama", "Alexeyka")
	vp.WriteConfig()

	// fsnotify
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("fileChange: %s\n", in.Name)
	})
	vp.WatchConfig()

}
