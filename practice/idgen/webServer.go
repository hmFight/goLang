package idgen

import (
	"fmt"
	"net/http"

	"strconv"

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
	http.HandleFunc("/id/incr/resetall", incrResetAll)
	http.HandleFunc("/id/incr/reset", incrReset)
	err := http.ListenAndServe(listenTo, nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func incrReset(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	paramKey, exist := request.Form["key"]
	var key string
	if !exist {
		fmt.Fprint(writer, "no param [key]")
	}
	key = paramKey[0]
	incrGen, existKey := concurrentMap.Get(key)
	if !existKey {
		fmt.Fprint(writer, "key:"+key+" not exist!")
	}
	gen, ok := incrGen.(*AutoIncrIdGen)
	if !ok {
		fmt.Fprint(writer, "error")
	}
	if gen.Reset() {
		fmt.Fprint(writer, "ok")
	} else {
		fmt.Fprint(writer, "failed")
	}
}

func incrResetAll(writer http.ResponseWriter, request *http.Request) {
	concurrentMap = initConnMap()
	fmt.Fprint(writer, "ok")
}

func incr(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	key := getParam(request, "key", "default")
	num, err := strconv.Atoi(getParam(request, "num", "1"))
	if err != nil {
		fmt.Fprint(writer, "param [num] error!")
	}
	tarIncrGen, mapExistKey := concurrentMap.Get(key)
	if !mapExistKey {
		tarIncrGen = NewAutoIncrIdGen()
		concurrentMap.Set(key, tarIncrGen)
	}
	idGen, ok := tarIncrGen.(IIdGenerator)
	if !ok {
		fmt.Fprint(writer, "unknown error!")
	}
	fmt.Fprint(writer, getIdRes(idGen, num))
}

//从http 表单中获取参数，如果不存在则返回默认
func getParam(request *http.Request, key string, defVal string) string {
	val, exist := request.Form[key]
	if !exist {
		return defVal
	} else {
		return val[0]
	}
}

func snowflake(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	num, err := strconv.Atoi(getParam(request, "num", "1"))
	if err != nil {
		fmt.Fprint(writer, "param [num] error!")
	}
	fmt.Fprint(writer, getIdRes(snowflakeIdGen, num))
}

func getIdRes(generator IIdGenerator, num int) string {
	var idResults string
	for i := 0; i < num; i++ {
		id := generator.GetId()
		if i > 0 {
			idResults += ","
		}
		idResults = idResults + strconv.FormatInt(id, 10)
	}
	return idResults
}
