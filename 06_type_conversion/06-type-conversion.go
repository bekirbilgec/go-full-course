package main

import "fmt"

func main() {

	/* x := 10
	y := 10.5

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion type(value) --> int(y) =10.0 =>10

	// fmt.Println(x + int(y))

	fmt.Println(y + float64(x)) */

	/* 	var x int8 = 125
	   	var y int16

	   	y = int16(x)

	   	fmt.Println(y) */

	/* var x int16 = 100
	var y int8

	y = int8(x)

	fmt.Println(y) */

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)

	fmt.Println()

	fmt.Printf("%v %T\n", str1, str1)

}
