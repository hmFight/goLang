package sync

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func readRlocal(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "read end")
}

func readLock(i int) {
	println(i, "read start")
	m.Lock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "read end")
}
