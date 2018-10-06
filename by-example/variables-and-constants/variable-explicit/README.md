# 2.1 Explicit Variable

Explicit Variable in Go. Here's the full source code.

#### main.go

```go
package main

import "fmt"

func main() {
	var anInt int   // = 10
	var aStr string // = "I am a string"

	anInt = 10
	aStr = "I am a string"

	fmt.Println("Integer Explicit Variable: ", anInt)
	fmt.Println("String Explicit Variable: ", aStr)
}

```

#### Project Structure

```bash
variable-explicit
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── output.variable-explicit.sample.txt
    │   │   ├── output.variable-explicit.txt
    │   │   └── variable-explicit
    │   ├── linux
    │   │   └── variable-explicit
    │   ├── mac
    │   │   └── variable-explicit
    │   └── win
    │       └── variable-explicit.exe
    └── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Integer Explicit Variable:  10
String Explicit Variable:  I am a string

```

#### Golang Playground

[Klik here to Run Explicit Variable on Golang Playground](https://play.golang.org/p/tpPkxdn5zb0)