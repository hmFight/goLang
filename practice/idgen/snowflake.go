package idgen

import (
	"sync"
	"sync/atomic"
	"time"
)

type SnowflakeIdGen struct {
	machineId     int64
	dataCenterId  int64
	idSequence    int64
	lastTimeStamp int64
	lock          *sync.RWMutex
}

var workIdShift = uint64(12)
var dataCenterIdShift = uint64(17)
var timestampLeftShift = uint64(22)
var sequenceMask = int64(4095)
var baseTimestamp = int64(1516460341000)

func NewIdGenerator(machineId, dataCenterId int64) *SnowflakeIdGen {
	nowTimestamp := nowTimestamp()
	generator := SnowflakeIdGen{
		machineId:     machineId,
		dataCenterId:  dataCenterId,
		idSequence:    0,
		lastTimeStamp: nowTimestamp,
		lock:          new(sync.RWMutex),
	}
	return &generator
}

func nowTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func nextTimestamp(lastTimestamp int64) int64 {
	timestamp := nowTimestamp()
	for timestamp < lastTimestamp {
		timestamp = nowTimestamp()
	}
	return timestamp
}

func (this *SnowflakeIdGen) GetId() int64 {
	timestamp := nowTimestamp()
	if timestamp == this.lastTimeStamp {
		atomic.AddInt64(&(this.idSequence), 1)
		this.idSequence = this.idSequence & sequenceMask
		if this.idSequence == 0 {
			timestamp = nextTimestamp(this.lastTimeStamp)
		}
	} else {
		this.idSequence = 0
	}
	this.lastTimeStamp = timestamp
	this.lock.Lock()
	newIdTmp := (timestamp - baseTimestamp) << timestampLeftShift
	newId := newIdTmp | (this.dataCenterId << dataCenterIdShift) | (this.machineId << workIdShift) | this.idSequence
	this.lock.Unlock()
	return newId
}
