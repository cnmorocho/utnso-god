package entities

type Mutex struct {
	Id         int
	ResourceId int
	State      int
	ThreadId   int
}
