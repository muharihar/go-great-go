package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// using hello-world-server, we need to run hello-world-server application first.
	resp, err := http.Get("http://127.0.0.1:8134/hello-world")

	// using mocky.io, we need an internet connection
	// resp, err := http.Get("http://www.mocky.io/v2/5bb8df303100005d003ed8c4")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
