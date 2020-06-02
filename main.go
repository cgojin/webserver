package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "server listen address")
	dir := flag.String("dir", "", "Static file directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	log.Printf("Serving listen on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
