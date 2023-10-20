package main

import "fmt"

func main() {

	/*	var (
		name      string = "mahmut"
		age       int    = 12
		isMarried bool   = false

		weight float32 = 42.3
		height int16   = 125
	) */

	/* var name, age, isMarried, weight, height = "mahmut", 12, true, 42.3, 125 */

	/* name, age, isMarried, weight, height := "mahmut", 12, true, 42.3, 125 */

	/* fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height) */

	// zero values string

	var name string

	// zero values numeric

	var age int
	var weight float32

	// zero values bool

	var isMarried bool

	fmt.Println(name) // --> string = "" return

	fmt.Println(age)    // --> int = 0 return
	fmt.Println(weight) // floot -- = 0 return

	fmt.Println(isMarried) // bool --> = false return

}
