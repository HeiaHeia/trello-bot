package main

import (
	"fmt"
)

func main() {

	config := LoadConfig("config.json")
	fmt.Println(config)
}
