package idgen

import (
	"fmt"
	"net/http"
)

var snowflakeIdGen SnowflakeIdGen = NewIdGenerator(1, 1)
var incrIdGen AutoIncrIdGen = NewAutoIncrIdGen()

func IdWebServer() {

	http.HandleFunc("/id/snowflake", snowflake)
	http.HandleFunc("/id/incr", incr)
	fmt.Println("start listen:127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}
func incr(writer http.ResponseWriter, request *http.Request) {
	id := incrIdGen.GetId()
	fmt.Fprint(writer, id)
}

func snowflake(writer http.ResponseWriter, request *http.Request) {
	id := snowflakeIdGen.GetId()
	fmt.Fprint(writer, id)
}
