package idgen

import (
	"fmt"
	"testing"
)

func TestWebServer(t *testing.T) {
	IdWebServer("7888")
}

func TestSnowflakeIdGen_GetId(t *testing.T) {
	stimestamp := nowTimestamp()
	generator := NewIdGenerator(1, 1)

	for i := 0; i <= 500; i++ {
		go func() {
			id := generator.GetId()
			fmt.Println(generator.idSequence, generator.lastTimeStamp)
			//fmt.Println(generator.lastTimeStamp)
			fmt.Println(id)
		}()
	}
	fmt.Println("cost:", nowTimestamp()-stimestamp)
}
