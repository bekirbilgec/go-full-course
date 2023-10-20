package main

import "fmt"

func main() {

	/* 	uint8       unsigned  8-bit integers (0 to 255)
	   	uint16      unsigned 16-bit integers (0 to 65535)
	   	uint32      unsigned 32-bit integers (0 to 4294967295)
	   	uint64      unsigned 64-bit integers (0 to 18446744073709551615)
	   	int8        signed  8-bit integers (-128 to 127)
	   	int16       signed 16-bit integers (-32768 to 32767)
	   	int32       signed 32-bit integers (-2147483648 to 2147483647)
	   	int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807) */

	/* float32     IEEE-754 32-bit floating-point numbers
	float64     IEEE-754 64-bit floating-point numbers
	complex64   complex numbers with float32 real and imaginary parts
	complex128  complex numbers with float64 real and imaginary parts */

	var name string = "bekir"
	var age int = 30
	var isMarried bool = true
	var weight float32 = 72.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

}
