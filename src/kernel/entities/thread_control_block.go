package entities

type ThreadControlBlock struct {
	ThreadId int
	Priority int
}

func (tcb *ThreadControlBlock) GetId() int {
	return tcb.ThreadId
}
