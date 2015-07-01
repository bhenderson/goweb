package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var listenFlag = flag.String("l", ":6123", "bind address")

func main() {
	flag.Parse()

	log.Printf("listening on %s\n", *listenFlag)
	log.Fatal(
		http.ListenAndServe(
			*listenFlag,
			http.FileServer(
				http.Dir(os.Getenv("PWD")),
			),
		),
	)
}
