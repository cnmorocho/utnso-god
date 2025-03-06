package services

import (
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

func (ks *KernelService) CreateProcess(process *entities.ProcessControlBlock) {
	ks.Kernel.Scheduler.NewProcess(process)
	request := RequestCreateProcessDTO{Size: process.Size}
	res, _ := ks.MemoryService.CreateProcess(request)
	if res.Code == "OK" {
		newThread := entities.NewThread(0)
		ks.Kernel.Scheduler.NewThread(newThread)
	}
}

func (ks *KernelService) DecodeProgram(codePath string) (utils.Queue[entities.Instruction], error) {
	return ks.Kernel.DecodeProgram(codePath)
}
