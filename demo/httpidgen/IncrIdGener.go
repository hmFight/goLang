package httpidgen

import "sync"

type IncrIdGene struct {
	step int
	id   int64
	lock *sync.RWMutex
}
