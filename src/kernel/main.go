package main

import (
	"cnmorocho/utnso-god/src/kernel/entities"
	"cnmorocho/utnso-god/src/kernel/services"
	"cnmorocho/utnso-god/src/utils"
	"fmt"
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
	intructions, _ := kernel.DecodeProgram(codePath)
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

	codePath := args[0]
	processSize, _ := strconv.Atoi(args[1])

	if err != nil {
		logger.Log("ERROR", "El tamaño del proceso no es un número válido")
		return
	}

	kernel := GetKernelInstance()
	InitializeFirstProcess(kernel, codePath, processSize)

	StartServer()
}
