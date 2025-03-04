package entities

import (
	"cnmorocho/utnso-god/src/utils"
)

type Kernel struct {
	CodePath          string
	NewProcesses      utils.Queue[Process]
	ReadyProcesses    utils.Queue[Process]
	ExecutedProcesses utils.Queue[Process]
	BlockedProcesses  utils.Queue[Process]
	ExitedProcesses   utils.Queue[Process]
}

func (k *Kernel) Schedule() {
	k.CreateProcess()

	for !k.NewProcesses.IsEmpty() {
		k.InitializeProcess()
	}
}

func (k *Kernel) CreateProcess() {
	pid := utils.GenerateId()
	newProcess := Process{Id: pid}
	k.NewProcesses.Enqueue(newProcess)
}

func (k *Kernel) InitializeProcess() {
	process, _ := k.NewProcesses.Dequeue()
	process.CreateThread()
	k.ReadyProcesses.Enqueue(process)
}
