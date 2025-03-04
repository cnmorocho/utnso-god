package main

import (
	"cnmorocho/utnso-god/src/utils"
	"fmt"
	"net/http"
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

	// Validar argumentos source_file y size
	if !argsAreValid(args) {
		logger.Log("ERROR", "Error en los argumentos")
		return
	}
	// Iniciar Comunicación con CPU y Memory
	_, err = http.Get("http://localhost:8080/cpu/health")
	if err != nil {
		logger.Log("ERROR", "Error al conectar con CPU")
		return
	}
	_, err = http.Get("http://localhost:8080/memory/health")
	if err != nil {
		logger.Log("ERROR", "Error al conectar con Memory")
		return
	}
	// Iniciar Servidor
	StartServer()

	// Planificar procesos
}
