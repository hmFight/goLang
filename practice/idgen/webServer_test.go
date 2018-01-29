package idgen

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

const incrApi = "http://localhost:7888/id/incr"
const snowflakeApi = "http://localhost:7888/id/snowflake"
const spKeyUrl = incrApi + "?key=key1"
const resetAllUrl = incrApi + "/resetall"
const resetUrl = incrApi + "/reset?key=key1"

func TestBatchGetIncrId(t *testing.T) {
	resetResult := httpGetBodyStr(resetAllUrl)
	assertEquals(t, "ok", resetResult, "resetResult")

	ids := httpGetBodyStr(incrApi + "?num=10")
	idArr := strings.Split(ids, ",")

	for index, id := range idArr {
		assertEquals(t, strconv.Itoa(index+1), id, "id")
	}
}

func TestBatchGetSnowflakeId(t *testing.T) {
	resetResult := httpGetBodyStr(resetAllUrl)
	assertEquals(t, "ok", resetResult, "resetResult")

	ids := httpGetBodyStr(snowflakeApi + "?num=1000")
	idArr := strings.Split(ids, ",")
	assertEquals(t, 1000, len(idArr), "id size")
}

func TestWebServer(t *testing.T) {
	resetResult := httpGetBodyStr(resetAllUrl)
	assertEquals(t, "ok", resetResult, "resetResult")

	for defaultKeyId, i := "", 1; i <= 100; i++ {
		defaultKeyId = httpGetBodyStr(incrApi)
		assertEquals(t, strconv.Itoa(i), defaultKeyId, "defaultKeyId")
	}

	for spkeyId, i := "", 1; i <= 1000; i++ {
		spkeyId = httpGetBodyStr(spKeyUrl)
		assertEquals(t, strconv.Itoa(i), spkeyId, "spkeyId")
	}

	//test reset
	resetResult = httpGetBodyStr(resetUrl)
	assertEquals(t, "ok", resetResult, "resetResult")

	for spkeyId, i := "", 1; i <= 1000; i++ {
		spkeyId = httpGetBodyStr(spKeyUrl)
		assertEquals(t, strconv.Itoa(i), spkeyId, "spkeyId")
	}

}

func assertEquals(t *testing.T, expect interface{}, actual interface{}, msg string) {
	if expect != actual {
		t.Error(msg, "expect:", expect, "actual:", actual)
	}
}

func httpGetBodyStr(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func TestSnowflakeIdGen_GetId(t *testing.T) {
	generator := NewIdGenerator(1, 1)
	id := generator.GetId()
	fmt.Println(generator.idSequence, generator.lastTimeStamp)
	fmt.Println(id)
}
