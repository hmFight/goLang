package idgen

import (
	"sync"
)

type AutoIncrIdGen struct {
	step int
	id   uint64
	lock *sync.RWMutex
}

func (this AutoIncrIdGen) GetId() uint64 {
	return this.id
}
