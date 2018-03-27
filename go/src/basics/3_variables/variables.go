// Package main is special. Defines a standalone executable, not a library.
// In this package, function main is where the program starts.
package main

import "fmt"

// this is an package level declaration
const Pi = 3.14

func main() {

	// Both these variables are local to function main.
	var radiusA = 5.0
	var CircleA_area = Pi * (radiusA * radiusA)

	fmt.Printf("The area of the circle is = %g \n", CircleA_area)

	// 2nd way to do it.
	fmt.Printf("2nd way: The area of the cirle is %g \n", circle_area(radiusA))

	// More on variables.
	// This is the general declaration for a variable
	// 	var var_name = expression
	// Either the type or the expression can be omitted, but not both.
	// If type is omitted, it is inferred from the expression.
	// If the expression is omitted, the variable is set to zero value which is
	// 0 for numbers, false for boolean, "" for strings and nil for references type and interfaces.
	// For arrays or struct, all their elements are set to zero.

	var s string
	fmt.Println("The value of s is: ", s)

	// Other ways of declaring variables.
	//	var a, b, c int   //all ints
	//  var d, e, f = 1, 10.0, "this is f"
	// Note: Unused vars will trigger errors.

	// Short local variable declarations
	// With functions, the folowing cam be used to declare and initialize local variable
	//    var_name := expression

	hello_string := "Hello there!"
	fmt.Println(hello_string)

	// note := is a declaration, = is an assignment.

}

// functions can be declared ahead of use.
func circle_area(r float64) float64 {

	return Pi * (r * r)
}
