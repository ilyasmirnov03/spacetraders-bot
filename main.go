package main

import (
	"fmt"
	"ilyasmirnov03/spacetraders-bot/src"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env_file, _ := godotenv.Read(".env")
	os.Setenv("TOKEN", env_file["TOKEN"])

	var command_input string
	fmt.Println("Choose command: ")
	fmt.Println("1. Mine")
	fmt.Println("2. Status")
	fmt.Scanln(&command_input)
	execute_command(command_input)
}

func execute_command(command string) {
	switch command {
	case "1":
		src.StartMining()
	case "2":
		get_status()
	}
}

func get_status() {
	body, err := src.CallApi[any]("", "GET", nil)
	if err != nil {
		return
	}
	fmt.Println(body)
}
