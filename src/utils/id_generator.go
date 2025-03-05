package utils

type IdGenerator struct {
	Id int
}

func NewIdGenerator() IdGenerator {
	return IdGenerator{
		Id: 0,
	}
}

func (ig *IdGenerator) New() int {
	ig.Id = ig.Id + 1
	return ig.Id
}

var ProcessIdGenerator = IdGenerator{Id: 0}
var ThreadIdGenerator = IdGenerator{Id: 0}
