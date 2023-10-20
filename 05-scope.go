package main

import "fmt"

// packVar := "Package Scope"   --> syntax error: short declaration non-declaration statement outside function body go

var packVar = "Package Scope"
var funcVar = "Func Scope(package)"

func main() {

	var funcVar = "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar)

}
