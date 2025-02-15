package main

import (
	"IAM/internal/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg.Server.Port)
}
