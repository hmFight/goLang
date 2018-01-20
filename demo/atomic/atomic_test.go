package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAdd(t *testing.T) {
	zero := int64(0)
	var id *int64 = &(zero)
	for i := 0; i <= 10; i++ {
		atomic.AddInt64(id, 1)
	}
	fmt.Println(*id)
}
