package idgen

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var help bool
var port int
var version bool
var apis string

func init() {
	flag.BoolVar(&version, "v", false, "version")
	flag.IntVar(&port, "port", 7888, "server port,default 7888")

	flag.StringVar(&apis, "apis", "",
		`/id/snowflake
	/id/incr`)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `hi,all you guys!
`)
		flag.PrintDefaults()
	}
}

func StartIdServer() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if version {
		fmt.Println("version:0.1.1")
		return
	}
	listenTo := ":" + strconv.Itoa(port)
	fmt.Println("listen:" + listenTo)
	fmt.Println("api:")
	fmt.Println("    /id/snowflake")
	fmt.Println("    /id/incr")

	IdWebServer(listenTo)
}
