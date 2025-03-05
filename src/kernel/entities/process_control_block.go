package entities

import (
	"cnmorocho/utnso-god/src/utils"
)

type ProcessControlBlock struct {
	ProcessId    int
	Size         int
	Instructions utils.Queue[Instruction]
	Mutex        Mutex
}

func (pcb *ProcessControlBlock) GetId() int {
	return pcb.ProcessId
}
