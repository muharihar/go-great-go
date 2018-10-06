# Hello World Main

Let's print "Hello, World!" using Go. Here's the full source code.

#### main.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

#### Project Structure

```bash
hello-world-main
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── hello-world-main
    │   │   ├── output.hello-world-main.sample.txt
    │   │   └── output.hello-world-main.txt
    │   ├── linux
    │   │   └── hello-world-main
    │   ├── mac
    │   │   └── hello-world-main
    │   └── win
    │       └── hello-world-main.exe
    └── main.go
```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Hello World!
```

#### Golang Playground

[Klik here to Run Hello World Main on Golang Playground](https://play.golang.org/p/oid3fhp9Cru)