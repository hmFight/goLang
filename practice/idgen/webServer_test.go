package idgen

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestWebServerId(t *testing.T) {
	prepare := make(chan bool)
	go func() {
		IdWebServer("localhost:7888", prepare)
	}()
	<-prepare
	funcName(t)
}

func TestWebServer(t *testing.T) {
	IdWebServer("localhost:7888", make(chan bool))
}
func funcName(t *testing.T) {
	defaultKeyId := httpGetBodyStr("localhost:7888/id/incr")
	key1Id := httpGetBodyStr("localhost:7888/id/incr?key=key1")
	if defaultKeyId != "1" {
		t.Error("expect 1,actual:", defaultKeyId)
	}
	if key1Id != "1" {
		t.Error("expect 1,actual:", key1Id)
	}
	helper(t, "1", defaultKeyId)
	helper(t, "1", key1Id)
	defaultKeyId = httpGetBodyStr("localhost:7888/id/incr")
	helper(t, "2", defaultKeyId)
	helper(t, "3", defaultKeyId)
}

func helper(t *testing.T, expect interface{}, actual interface{}) {
	if expect != actual {
		t.Error("expect:", expect, "actual:actual")
	}
}
func httpGetBodyStr(url string) string {
	resp, err := http.Get("")
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
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
