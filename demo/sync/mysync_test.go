package sync

import (
	"sync"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	m = new(sync.RWMutex)
	for i := 0; i < 10; i++ {
		go readRlocal(i)
		//go read(i)
	}
	time.Sleep(2 * time.Second)
}

func TestReadLock(t *testing.T) {
	m = new(sync.RWMutex)
	for i := 0; i < 10; i++ {
		go readLock(i)
	}
	time.Sleep(100 * time.Second)
}
