package entities

import "cnmorocho/utnso-god/src/utils"

type Mutex struct {
	IsLocked bool
	Owner    *ThreadControlBlock
	Queue    *utils.Queue[ThreadControlBlock]
}
