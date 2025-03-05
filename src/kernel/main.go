package main

import (
	"cnmorocho/utnso-god/src/utils"
	"fmt"
	"os"
)

func argsAreValid(args []string) bool {
	if len(args) != 2 {
		return false
	}

	if args[0] == "" || args[1] == "" {
		return false
	}

	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	args := os.Args[1:]

	err := utils.NewLogger("kernel.log")

	if err != nil {
		fmt.Println("Error al iniciar logs")
		return
	}

	logger := utils.GetLogger()

	if !argsAreValid(args) {
		logger.Log("ERROR", "Error en los argumentos")
		return
	}

	codePath := os.Args[0]
	processSize := os.Args[1]

	StartServer()
}
