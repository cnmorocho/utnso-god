package main

import (
	"cnmorocho/utnso-god/src/kernel/dtos"
	"cnmorocho/utnso-god/src/kernel/entities"
	"cnmorocho/utnso-god/src/kernel/services"
	"cnmorocho/utnso-god/src/utils"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

var kernelInstance *entities.Kernel

func GetKernelInstance() *entities.Kernel {
	if kernelInstance == nil {
		kernelInstance = &entities.Kernel{
			Scheduler: entities.NewScheduler(),
		}
	}
	return kernelInstance
}

func InitializeFirstProcess(firstProcess dtos.CreateProcessRequestDTO) {
	kernel := GetKernelInstance()
	memoryService := services.NewMemoryService()
	kernelService := services.NewKernelService(kernel, memoryService)
	kernelService.CreateProcess(firstProcess)
}

func validateArguments(args []string) bool {
	if len(args) != 2 {
		return false
	}

	if args[0] == "" || args[1] == "" {
		slog.Error("no se ingresaron argumentos")
		return false
	}

	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		slog.Error(fmt.Sprintf("no se encontro el archivo %s", args[0]))
		return false
	}

	_, err := strconv.Atoi(args[1])
	if err != nil {
		slog.Error("el tamaño del proceso no es un número válido")
		return false
	}

	return true
}

func main() {
	args := os.Args[1:]

	utils.InitLogger("kernel.log")
	validateArguments(args)

	codePath := args[0]
	processSize, _ := strconv.Atoi(args[1])

	InitializeFirstProcess(dtos.CreateProcessRequestDTO{FilePath: codePath, ProcessSize: processSize, ThreadPriority: 0})
	StartServer()
}
