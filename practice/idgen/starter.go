package idgen

import (
	"flag"
	"fmt"
	"strconv"
)

var port int
var version bool
var apis bool

func init() {
	flag.BoolVar(&version, "v", false, "show version")
	flag.IntVar(&port, "port", 7888, "server port,default 7888")

	flag.BoolVar(&apis, "apis", false, "how to use")

	//	flag.Usage = func() {
	//		fmt.Fprintf(os.Stdout, `hi,all you guys!
	//`)
	//		flag.PrintDefaults()
	//	}
}

func StartIdServer() {
	flag.Parse()
	if version {
		fmt.Println("version:0.5.1")
		return
	}
	if apis {
		fmt.Println(
			"/id/snowflake\n" +
				"    --get id by snowfalke\n" +
				"/id/incr[?key=your_key]\n" +
				"    --get id by auto-incrementor,every key map a IdGenerator\n" +
				"/id/incr/resetall\n" +
				"    --reset all AutoIdGenerator to 0\n" +
				"/id/incr/reset[?key=your_key]\n" +
				"    --reset the IdGenerator of your_key to 0")
		return
	}
	listenTo := ":" + strconv.Itoa(port)
	fmt.Println("listen" + listenTo)

	IdWebServer(listenTo)
}
