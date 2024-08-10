package main

import (
	"fmt"
	"github.com/zemartins81/posgoApi/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config)

}
