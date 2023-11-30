package main

import (
	"fmt"
	"project-api/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	envs := config.GetEnvs()

	fmt.Println(envs)
}
