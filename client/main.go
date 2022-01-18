package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func usage() {
	name := filepath.Base(os.Args[0])
	fmt.Printf("Usage:\n  %s [options] <URL>\nExamles:\n"+
		"  %s http://localhost:8080\n"+
		"  %s https://localhost:8080 -insecure\n"+
		"  %s https://www.google.com\n"+
		"Options:\n",
		name, name, name, name)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	insecure := flag.Bool("insecure", false, "accepts any certificate, only for testing.")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
		return
	}
	url := args[0]
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// curl with -k (or –insecure) option, example: curl -k https://localhost:8080
				InsecureSkipVerify: *insecure,
			},
			Proxy: http.ProxyFromEnvironment,
		},
	}

	//resp, err := http.Get(url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Failed to request URL(%s): %v\n", url, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response(%s): %v\n", url, err)
		return
	}

	fmt.Println(string(body))
}
