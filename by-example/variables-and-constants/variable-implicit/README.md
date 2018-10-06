# Implicit Variable

Implicit Variable in Go. Here's the full source code.

#### main.go

```go
package main

import "fmt"

func main() {
	// var anInt int   // = 10
	// var aStr string // = "I am a string"

	anInt := 10
	aStr := "I am a string"

	fmt.Println("Integer Implicit Variable: ", anInt)
	fmt.Println("String Implicit Variable: ", aStr)
}

```

#### Project Structure

```bash
variable-implicit
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── output.variable-implicit.sample.txt
    │   │   ├── output.variable-implicit.txt
    │   │   └── variable-implicit
    │   ├── linux
    │   │   └── variable-implicit
    │   ├── mac
    │   │   └── variable-implicit
    │   └── win
    │       └── variable-implicit.exe
    └── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Integer Implicit Variable:  10
String Implicit Variable:  I am a string

```

#### Golang Playground

[Klik here to Run Implicit Variable on Golang Playground](https://play.golang.org/p/fI7tsQBVGV2)