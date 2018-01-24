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

func (this *AutoIncrIdGen) Reset() bool {
	oldVal := *this.id
	swappedOk := atomic.CompareAndSwapInt64(this.id, oldVal, 0)
	return swappedOk
}

func NewAutoIncrIdGen() AutoIncrIdGen {
	zero := int64(0)
	gen := AutoIncrIdGen{
		id:   &zero,
		step: 1,
		lock: new(sync.RWMutex),
	}
	return gen
}
