package entities

import "cnmorocho/utnso-god/src/utils"

type Process struct {
	Id              int
	Size            int
	NewThreads      utils.Queue[Thread]
	ReadyThreads    utils.Queue[Thread]
	ExecutedThreads utils.Queue[Thread]
	BlockedThreads  utils.Queue[Thread]
	ExitedThreads   utils.Queue[Thread]
}

func (p *Process) CreateThread() {
	tid := utils.GenerateId()
	p.NewThreads.Enqueue(Thread{Id: tid})
}
