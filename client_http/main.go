package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// similar to: curl http://localhost:8080
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
