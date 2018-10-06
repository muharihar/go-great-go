# Implicit Constant

Implicit Constant in Go. Here's the full source code.

#### main.go

```go
package main

import "fmt"

func main() {
	const anInt = 10
	const aStr = "I am a string"

	fmt.Println("Integer Implicit Constant: ", anInt)
	fmt.Println("String Implicit Constant: ", aStr)
}

```

#### Project Structure

```bash
constant-implicit
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── constant-implicit
    │   │   ├── output.constant-implicit.sample.txt
    │   │   └── output.constant-implicit.txt
    │   ├── linux
    │   │   └── constant-implicit
    │   ├── mac
    │   │   └── constant-implicit
    │   └── win
    │       └── constant-implicit.exe
    └── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Integer Implicit Constant:  10
String Implicit Constant:  I am a string

```

#### Golang Playground

[Klik here to Run Implicit Constant on Golang Playground](https://play.golang.org/p/mhijniabRhJ)