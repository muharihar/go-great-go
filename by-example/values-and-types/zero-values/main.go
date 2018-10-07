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
