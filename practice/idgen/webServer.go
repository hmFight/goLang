package idgen

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var help bool
var port int
var version bool
var apis string
var snowflakeIdGen = NewIdGenerator(1, 1)
var incrIdGen = NewAutoIncrIdGen()

func init() {
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&help, "help", false, "this help")
	flag.IntVar(&port, "port", 7888, "server port,default 7888")

	flag.StringVar(&apis, "apis", "",
		`/id/snowflake
	/id/incr`)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `hi,all you guys!
`)
		flag.PrintDefaults()
	}
}

func IdWebServer() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if version {
		fmt.Println("version:0.1.1")
		return
	}
	http.HandleFunc("/id/snowflake", snowflake)
	http.HandleFunc("/id/incr", incr)
	listenTo := "127.0.0.1:" + strconv.Itoa(port)
	fmt.Println("listen:" + listenTo)
	fmt.Println("api:")
	fmt.Println("    /id/snowflake")
	fmt.Println("    /id/incr")
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
