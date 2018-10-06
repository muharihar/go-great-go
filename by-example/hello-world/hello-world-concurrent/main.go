package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello World! #1")
	}()

	go func() {
		fmt.Println("Hello World! #2")
	}()

	go func() {
		fmt.Println("Hello World! #3")
	}()

	time.Sleep(100 * time.Millisecond)
}
