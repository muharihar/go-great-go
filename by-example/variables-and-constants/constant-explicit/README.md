# Explicit Constant

Explicit Constant in Go. Here's the full source code.

#### main.go

```go
package main

import (
	"fmt"
)

func main() {
	const anInt int = 10
	const aStr string = "I am a string"

	fmt.Println("Integer Explicit Constant: ", anInt)
	fmt.Println("String Explicit Constant: ", aStr)
}

```

#### Project Structure

```bash
constant-explicit
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── constant-explicit
    │   │   ├── output.constant-explicit.sample.txt
    │   │   └── output.constant-explicit.txt
    │   ├── linux
    │   │   └── constant-explicit
    │   ├── mac
    │   │   └── constant-explicit
    │   └── win
    │       └── constant-explicit.exe
    └── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Integer Explicit Constant:  10
String Explicit Constant:  I am a string

```

#### Golang Playground

[Klik here to Run Explicit Constant on Golang Playground](https://play.golang.org/p/d7qtBGw8CUy)