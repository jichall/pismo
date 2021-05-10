package main

import "flag"

var (
	host = flag.String("host", "127.0.0.1", "")
	port = flag.String("port", "8090", "")
)

func main() {
	flag.Parse()

	serve(*host, *port)
}
