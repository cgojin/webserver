package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// command params
	addr := flag.String("addr", ":8080", "server listen address")
	dir := flag.String("dir", "", "static file directory")
	certFile := flag.String("cert", "", "certificate file")
	keyFile := flag.String("key", "", "private key file")
	flag.Parse()

	if len(*addr) <= 0 {
		log.Fatal("no listen address")
	}

	// print url
	var err error
	var url string
	useTLS := len(*certFile) > 0 && len(*keyFile) > 0
	if useTLS {
		url = "https://"
	} else {
		url = "http://"
	}
	if (*addr)[0] == ':' {
		url = url + "localhost" + *addr
	} else {
		url = url + *addr
	}
	fmt.Printf("web server: %s", url)

	// start server
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	if useTLS {
		err = http.ListenAndServeTLS(*addr, *certFile, *keyFile, nil)
	} else {
		err = http.ListenAndServe(*addr, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}
