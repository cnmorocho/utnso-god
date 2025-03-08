package main

import (
	"cnmorocho/utnso-god/src/kernel/entities"
	"cnmorocho/utnso-god/src/kernel/services"
	"cnmorocho/utnso-god/src/utils"
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

func InitializeFirstProcess(kernel *entities.Kernel, codePath string, processSize int) {
	intructions, _ := entities.DecodeProgram(codePath)
	newProcess := entities.ProcessControlBlock{
		ProcessId:    0,
		Size:         processSize,
		Instructions: intructions,
		Mutex:        entities.Mutex{},
	}
	memoryService := services.NewMemoryService()
	kernelService := services.NewKernelService(kernel, memoryService)
	kernelService.CreateProcess(&newProcess)
}

func validateArguments(args []string) bool {
	if len(args) != 2 {
		return false
	}

	if args[0] == "" || args[1] == "" {
		slog.Error("se deben ingresar argumentos")
		return false
	}

	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		slog.Error("se deben ingresar argumentos")
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

	kernel := GetKernelInstance()
	InitializeFirstProcess(kernel, codePath, processSize)

	StartServer()
}
