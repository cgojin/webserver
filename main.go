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
	dir := flag.String("dir", "", "Static file directory")
	certFile := flag.String("cert", "", "Certificate file")
	keyFile := flag.String("key", "", "Private key file")
	flag.Parse()

	if len(*addr) <= 0 {
		log.Fatal("No listen address")
	}

	// print url
	var err error
	var prefix string
	useTLS := len(*certFile) > 0 && len(*keyFile) > 0
	if useTLS {
		prefix = "https://"
	} else {
		prefix = "http://"
	}
	fmt.Printf("Web Server: ")
	if (*addr)[0] == ':' {
		fmt.Printf("%s", prefix+"localhost"+*addr)
	} else {
		fmt.Printf("%s\n", prefix+*addr)
	}

	// start server
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	if useTLS {
		err = http.ListenAndServeTLS(":8080", *certFile, *keyFile, nil)
	} else {
		err = http.ListenAndServe(*addr, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}
