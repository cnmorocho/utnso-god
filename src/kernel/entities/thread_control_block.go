package entities

import "cnmorocho/utnso-god/src/utils"

type ThreadControlBlock struct {
	ThreadId int
	Priority int
}

func NewThread(priority int) *ThreadControlBlock {
	return &ThreadControlBlock{
		ThreadId: utils.ThreadIdGenerator.Id,
		Priority: priority,
	}
}

func (tcb *ThreadControlBlock) GetId() int {
	return tcb.ThreadId
}
