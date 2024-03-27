package main

import (
	"marketplace/app/routers"
	"fmt"
)

func main() {
	err := routers.Run()
	if err != nil {
		fmt.Println(err)
	}
}
