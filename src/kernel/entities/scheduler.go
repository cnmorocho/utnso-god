package entities

import "cnmorocho/utnso-god/src/utils"

type Schedulable interface {
	GetId() int
}

type Scheduler struct {
	NewProcesses      *utils.Queue[ProcessControlBlock]
	ReadyProcesses    *utils.Queue[ProcessControlBlock]
	ExecutedProcesses *utils.Queue[ProcessControlBlock]
	BlockedProcesses  *utils.Queue[ProcessControlBlock]
	ExitedProcesses   *utils.Queue[ProcessControlBlock]
	NewThreads        *utils.Queue[ThreadControlBlock]
	ReadyThreads      *utils.Queue[ThreadControlBlock]
	ExecutedThreads   *utils.Queue[ThreadControlBlock]
	BlockedThreads    *utils.Queue[ThreadControlBlock]
	ExitedThreads     *utils.Queue[ThreadControlBlock]
}

func (s *Scheduler) MoveToQueue(from *utils.Queue[Schedulable], to *utils.Queue[Schedulable]) {
	elem, _ := from.Dequeue()
	to.Enqueue(elem)
}

func (s *Scheduler) NewProcess(elem ProcessControlBlock) {
	s.NewProcesses.Enqueue(elem)
}

func (s *Scheduler) NewThread(elem ThreadControlBlock) {
	s.NewThreads.Enqueue(elem)
}
