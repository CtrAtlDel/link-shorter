package main

import (
	"fmt"
	"ik/linkShorter/internal/config"
)

func main() {
	cfg:= config.MustLoad()

	fmt.Println(cfg)
}

func setupLogger(env string) {
	
}