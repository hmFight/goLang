package idgen

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestWebServerId(a *testing.T) {
	IdWebServer(":7888")
}

func TestWebServer(t *testing.T) {
	defaultKeyIdUrl := "http://localhost:7888/id/incr"
	spKeyUrl := "http://localhost:7888/id/incr?key=key1"

	resetResult := httpGetBodyStr("http://localhost:7888/id/incr/reset")
	testHelper(t, "ok", resetResult, "resetResult")

	for defaultKeyId, i := "", 1; i <= 100; i++ {
		defaultKeyId = httpGetBodyStr(defaultKeyIdUrl)
		testHelper(t, strconv.Itoa(i), defaultKeyId, "defaultKeyId")
	}

	for spkeyId, i := "", 1; i <= 1000; i++ {
		spkeyId = httpGetBodyStr(spKeyUrl)
		testHelper(t, strconv.Itoa(i), spkeyId, "spkeyId")
	}

	resetResult = httpGetBodyStr("http://localhost:7888/id/incr/reset")
	testHelper(t, "ok", resetResult, "resetResult")

}

func testHelper(t *testing.T, expect interface{}, actual interface{}, msg string) {
	if expect != actual {
		t.Error(msg, "expect:", expect, "actual:", actual)
	}
}
func httpGetBodyStr(url string) string {
	resp, err := http.Get(url)
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
			fmt.Println(id)
		}()
	}
	fmt.Println("cost:", nowTimestamp()-stimestamp)
}
