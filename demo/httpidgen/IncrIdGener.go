package httpidgen

import "sync"

type IncrIdGene struct {
	step int
	id   int
	lock *sync.RWMutex
}
