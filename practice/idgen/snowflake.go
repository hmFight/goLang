package idgen

import (
	"sync"
	"sync/atomic"
	"time"
)

type SnowflakeIdGen struct {
	MechineId    uint64
	DataCenterId uint64
	idSequence   *int64
	timeStamp    uint64
	lock         *sync.RWMutex
}

var workIdShift = uint64(12)
var dataCenterIdShift = uint64(17)
var timestampLeftShift = uint64(22)
var sequnceMask = uint64(4095)
var baseTimestamp = uint64(1318323746000)

func NewIdGenerator(mechineId, dataCenterId uint64) SnowflakeIdGen {
	timestamp := time.Now().UnixNano() / 1000000
	zero := int64(0)
	generator := SnowflakeIdGen{
		MechineId:    mechineId,
		DataCenterId: dataCenterId,
		idSequence:   &zero,
		timeStamp:    uint64(timestamp),
		lock:         new(sync.RWMutex),
	}
	go generator.cleanSequence()
	return generator
}

func (this SnowflakeIdGen) cleanSequence() {
	for {
		this.lock.Lock()
		atomic.SwapInt64(this.idSequence, int64(0))
		this.timeStamp = uint64(time.Now().UnixNano() / 1000000)
		this.lock.Unlock()
		time.Sleep(time.Second)
	}
}

func (this *SnowflakeIdGen) GetId() uint64 {
	seq := atomic.AddInt64(this.idSequence, 1)
	this.lock.RLock()
	newId := ((this.timeStamp - baseTimestamp) << timestampLeftShift) | (this.DataCenterId << dataCenterIdShift) | (this.MechineId << workIdShift) | uint64(seq)
	this.lock.RUnlock()
	return newId
}
