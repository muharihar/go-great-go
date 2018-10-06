# 1.2. Hello World Concurrent

Let's use Go to print "Hello World!" in Concurrent mode. In Go, you need to use goroutine to represent Concurrent  of execution.

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

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Hello World! #2
Hello World! #3
Hello World! #1

```

#### Running with Golang Playground

[Klik here to Run Hello World Concurrent on Golang Playground](https://play.golang.org/p/s7Gv1MuVcSp)