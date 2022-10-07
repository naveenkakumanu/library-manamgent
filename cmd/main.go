package main

import (
	"flag"
	"os"

	router "github.com/naveenkakumanu/go-crud-gin/routers"
)

var port string

var ports = flag.String("port", ":8081", "Port to Execute the application")

// Reading Command Line Arguments with help of os and flag package

func main() {

	flag.Parse()

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	if port == "" {
		port = *ports
	}

	router := router.Router()
	router.Run(port)
}