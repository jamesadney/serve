package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	cache := flag.Bool("c", false, "enable caching")

	flag.Parse()

	directory := flag.Arg(0)
	dirPath, err := filepath.Abs(directory)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%d", *port)

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
		ip = "127.0.0.1"
	}

	fmt.Printf("Serving '%s/' on 'http://%s%s'\n", dirPath, ip, address)
	log.Fatal(ServeDir(dirPath, address, *cache))
}
