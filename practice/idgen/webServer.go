package idgen

import (
	"fmt"
	"net/http"
)

var snowflakeIdGen = NewIdGenerator(1, 1)
var incrIdGen = NewAutoIncrIdGen()

func IdWebServer(listenTo string) {
	http.HandleFunc("/id/snowflake", snowflake)
	http.HandleFunc("/id/incr", incr)
	err := http.ListenAndServe(listenTo, nil)

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
