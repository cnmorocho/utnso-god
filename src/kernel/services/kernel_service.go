package services

import (
	"cnmorocho/utnso-god/src/kernel/dtos"
	"cnmorocho/utnso-god/src/kernel/entities"
	"cnmorocho/utnso-god/src/utils"
)

type KernelService struct {
	Kernel        *entities.Kernel
	MemoryService *MemoryService
}

func NewKernelService(kernel *entities.Kernel, memoryService *MemoryService) *KernelService {
	return &KernelService{
		Kernel:        kernel,
		MemoryService: memoryService,
	}
}

func (ks *KernelService) CreateProcess(newProcess dtos.CreateProcessRequestDTO) {
	id := utils.ProcessIdGenerator.New()
	decodedInstructions, _ := entities.DecodeProgram(newProcess.FilePath)

	newPcb := entities.ProcessControlBlock{
		ProcessId:    id,
		Size:         newProcess.ProcessSize,
		Instructions: decodedInstructions,
		Mutex:        entities.Mutex{},
	}
	ks.Kernel.Scheduler.NewProcess(&newPcb)
	request := RequestCreateProcessDTO{Size: newProcess.ProcessSize}
	res, _ := ks.MemoryService.CreateProcess(request)
	if res.Code == "OK" {
		newThread := entities.NewThread(0)
		ks.Kernel.Scheduler.NewThread(newThread)
	}
}
