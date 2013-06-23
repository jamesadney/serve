package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	port int
	host string
)

func init() {
	const (
		DEFAULT_PORT = 8000
		PORT_USAGE   = "port to run server on"
	)
	flag.IntVar(&port, "port", DEFAULT_PORT, PORT_USAGE)
	flag.IntVar(&port, "p", DEFAULT_PORT, PORT_USAGE)

	const (
		DEFAULT_HOST = "127.0.0.1"
		HOST_USAGE   = "host to run server on"
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, HOST_USAGE)
	flag.StringVar(&host, "h", DEFAULT_HOST, HOST_USAGE)
}

func main() {
	flag.Parse()

	directory := flag.Arg(0)
	absPath, err := filepath.Abs(directory)
	if err != nil {
		panic(err)
	}

	location := host + ":" + strconv.Itoa(port)
	fmt.Printf("Serving '%s/' on 'http://%s'\n", absPath, location)
	panic(http.ListenAndServe(location, http.FileServer(http.Dir(absPath))))
}
