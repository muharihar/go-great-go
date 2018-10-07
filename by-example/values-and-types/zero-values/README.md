# 3.1. Zero Values

Basic zero values in Go. Here's the full source code.

#### main.go

```go
package main

import (
	"fmt"
)

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Printf("i (%T) : %v \n", i, i)
	fmt.Printf("i8 (%T) : %v \n", i8, i8)
	fmt.Printf("i16 (%T) : %v \n", i16, i16)
	fmt.Printf("i32 (%T) : %v \n", i32, i32)
	fmt.Printf("i64 (%T) : %v \n", i64, i64)

	var f32 float32
	var f64 float64

	fmt.Printf("f32 (%T) : %v \n", f32, f32)
	fmt.Printf("f64 (%T) : %v \n", f64, f64)

	var b bool
	var s string

	fmt.Printf("b (%T) : %v \n", b, b)
	fmt.Printf("s (%T) : %q \n", s, s)
}

```

#### Project Structure

```bash
zero-values
	├── README.md
	├── build_and_run.sh
	├── build_and_run_with_output.sh
	├── dist
	│   ├── default
	│   │   ├── output.zero-values.sample.txt
	│   │   ├── output.zero-values.txt
	│   │   └── zero-values
	│   ├── linux
	│   │   └── zero-values
	│   ├── mac
	│   │   └── zero-values
	│   └── win
	│       └── zero-values.exe
	└── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
i (int) : 0
i8 (int8) : 0
i16 (int16) : 0
i32 (int32) : 0
i64 (int64) : 0
f32 (float32) : 0
f64 (float64) : 0
b (bool) : false
s (string) : ""

```

#### Golang Playground

[Klik here to Run Zero Values on Golang Playground](https://play.golang.org/p/x0Av6YkkWwd)