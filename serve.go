package main

import (
	"flag"
	"fmt"
	"github.com/jamesadney/serve/fileserver"
	"log"
	"path/filepath"
	"strconv"
)

var (
	host  string
	port  int
	cache bool
)

func init() {
	const (
		DEFAULT_HOST = "127.0.0.1"
		HOST_USAGE   = "host to run server on"
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, HOST_USAGE)
	flag.StringVar(&host, "h", DEFAULT_HOST, HOST_USAGE)

	const (
		DEFAULT_PORT = 8000
		PORT_USAGE   = "port to run server on"
	)
	flag.IntVar(&port, "port", DEFAULT_PORT, PORT_USAGE)
	flag.IntVar(&port, "p", DEFAULT_PORT, PORT_USAGE)

	const (
		DEFAULT_CACHE = false
		CACHE_USAGE   = "enable caching"
	)
	flag.BoolVar(&cache, "enable-caching", DEFAULT_CACHE, CACHE_USAGE)
	flag.BoolVar(&cache, "c", DEFAULT_CACHE, CACHE_USAGE)
}

func main() {
	flag.Parse()

	directory := flag.Arg(0)
	dirPath, err := filepath.Abs(directory)
	if err != nil {
		panic(err)
	}

	address := host + ":" + strconv.Itoa(port)

	fmt.Printf("Serving '%s/' on 'http://%s'\n", dirPath, address)
	log.Fatal(fileserver.ServeDir(dirPath, address, cache))
}
