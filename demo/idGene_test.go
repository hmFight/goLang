package demo

import (
	"fmt"
	"github.com/fanngyuan/idgen"
	"testing"
)

// id 生成
func TestIdGene(t *testing.T) {
	generator1 := idgen.NewIdGenerator(1, 1)
	generator2 := idgen.NewIdGenerator(2, 1)
	generator3 := idgen.NewIdGenerator(3, 1)
	for i := 0; i < 100; i++ {
		fmt.Printf("generator1(dc %d,machine %d):%d \n", generator1.DataCenterId, generator1.MechineId, generator1.GetId())
		fmt.Printf("generator2(dc %d,machine %d):%d \n", generator2.DataCenterId, generator2.MechineId, generator2.GetId())
		fmt.Printf("generator3(dc %d,machine %d):%d \n", generator3.DataCenterId, generator3.MechineId, generator3.GetId())
		fmt.Println(i)
	}
}
