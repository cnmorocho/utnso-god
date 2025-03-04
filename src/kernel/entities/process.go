package entities

import "cnmorocho/utnso-god/src/utils"

type Process struct {
	Id             int
	Size           int
	NewTreads      utils.Queue[Thread]
	ReadyTreads    utils.Queue[Thread]
	ExecutedTreads utils.Queue[Thread]
	BlockedTreads  utils.Queue[Thread]
	ExitedTreads   utils.Queue[Thread]
}

func (p *Process) CreateThread() {
	tid := utils.GenerateId()
	p.NewTreads.Enqueue(Thread{Id: tid})
}
