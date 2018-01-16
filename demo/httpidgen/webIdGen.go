package httpidgen

import (
	"fmt"
	"net/http"
	"strconv"
)

var id int = 0

func GenIncrId() {
	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/id", idgen)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("127.0.0.1:7878", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func idgen(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, incr(1))
}

func incr(step int) string {
	id += step
	return strconv.Itoa(id)

}
