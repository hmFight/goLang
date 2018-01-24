package idgen

import (
	"fmt"
	"net/http"

	"github.com/orcaman/concurrent-map"
)

var snowflakeIdGen = NewIdGenerator(1, 1)
var concurrentMap = initConnMap()

func initConnMap() cmap.ConcurrentMap {
	connMap := cmap.New()
	connMap.Set("default", NewAutoIncrIdGen())
	return connMap
}

func IdWebServer(listenTo string) {
	http.HandleFunc("/id/snowflake", snowflake)
	http.HandleFunc("/id/incr", incr)
	http.HandleFunc("/id/incr/reset", incrReset)
	err := http.ListenAndServe(listenTo, nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func incrReset(writer http.ResponseWriter, request *http.Request) {
	concurrentMap = initConnMap()
	fmt.Fprint(writer, "ok")
}

func incr(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	paramKey, exist := request.Form["key"]
	var key string
	if !exist {
		key = "default"
	} else {
		key = paramKey[0]
	}
	tarIncrGen, mapExistKey := concurrentMap.Get(key)
	if !mapExistKey {
		tarIncrGen = NewAutoIncrIdGen()
		concurrentMap.Set(key, tarIncrGen)
	}
	gen := tarIncrGen.(AutoIncrIdGen)
	fmt.Fprint(writer, gen.GetId())
}

func snowflake(writer http.ResponseWriter, request *http.Request) {
	id := snowflakeIdGen.GetId()
	fmt.Fprint(writer, id)
}
