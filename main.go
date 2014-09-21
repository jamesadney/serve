package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"log"
	"path/filepath"
)

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	cache := flag.Bool("c", false, "enable caching")
	browser := flag.Bool("b", false, "open browser")

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
		ip = "0.0.0.0"
	}

	url := fmt.Sprintf("http://%s%s", ip, address)
	fmt.Printf("Serving '%s/' on '%s'\n", dirPath, url)
	if *browser {
		open.Run(url)
	}
	// Blocks here
	log.Fatal(ServeDir(dirPath, address, *cache))
}
