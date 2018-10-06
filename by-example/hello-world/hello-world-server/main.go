package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	})

	fmt.Println("Press ctrl+c to shutdown the application.")
	log.Fatal(http.ListenAndServe("127.0.0.1:8134", nil))
}
