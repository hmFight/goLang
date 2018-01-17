package idgen

import (
	"sync"
	"sync/atomic"
)

type AutoIncrIdGen struct {
	step int64
	id   *int64
	lock *sync.RWMutex
}

func (this *AutoIncrIdGen) GetId() uint64 {
	res := atomic.AddInt64(this.id, this.step)
	return uint64(res)
}
