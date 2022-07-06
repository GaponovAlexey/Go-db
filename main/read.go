package main

import (
	"conect-db/ut"
	"fmt"
)

func main() {

	data, err := ut.GetConfigJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

}
