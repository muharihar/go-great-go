# Hello World Concurrent
Let's use Go to print "Hello, World!" in Concurrent. In Go, you need to use goroutine to represent Concurrent  of execution.


#### Golang Playground

[Klik here to Run Hello World Concurrent on Golang Playground](https://play.golang.org/p/s7Gv1MuVcSp)

#### main.go

```go
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
```

#### Project Structure

```bash
hello-world-concurrent
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── hello-world-concurrent
    │   │   ├── output.hello-world-concurrent.sample.txt
    │   │   └── output.hello-world-concurrent.txt
    │   ├── linux
    │   │   └── hello-world-concurrent
    │   ├── mac
    │   │   └── hello-world-concurrent
    │   └── win
    │       └── hello-world-concurrent.exe
    └── main.go
```