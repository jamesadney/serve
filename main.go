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

	fmt.Printf("Serving '%s/' on 'http://%s'\n", dirPath, address)
	log.Fatal(ServeDir(dirPath, address, *cache))
}
