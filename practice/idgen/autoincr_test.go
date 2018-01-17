package idgen

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestIncr(t *testing.T) {
	zero := int64(0)
	gen := AutoIncrIdGen{
		id:   &zero,
		step: 1,
		lock: new(sync.RWMutex),
	}
	for i := 0; i < 100; i++ {
		go fmt.Println(gen.GetId())
	}
	time.Sleep(5 * time.Second)
	fmt.Println("last:", gen.GetId())
}
