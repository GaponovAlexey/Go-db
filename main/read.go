package main

import (
	"conect-db/utils"
	"fmt"
)

func main() {

	data, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

}
