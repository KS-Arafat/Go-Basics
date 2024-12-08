// Source: https://go.dev/tour
package main

// NB: import(...) for importing more than one packages
import (
	"fmt"
)

// Const has to be declared by '=' not this':=', Can be local or global
// NB: const(...) for declaring more than one constants
const GLOBAL_CONST = 5

func main() {
	/////////////////
	// Chapter 01: Basics
	/////////////////
	i, j := Div(6, 4)
	// Like C printf with different types of flags
	fmt.Printf("Result %T : %v, %T : %v\n", i, i, j, j)
	a := ToFloat(6)
	fmt.Printf("Result %T : %v\n", a, a)

	///////////////
	// Chapter 02: Flow Control
	///////////////

	sum := 0
	// Like for loop in C but dont need paranthesis
	// also with hard syntax with curly braces like functions
	for i := 0; i < 5; i++ {
		sum += i
	}
	// for also acts like while loop in C
	for sum < 50 {
		sum += 1
	}
	// infinte loop
	for {
		sum += 1
		if sum%11 == 0 {
			break
		}
	}

	// like if in C but without the paranthesis
	// also with hard syntax with curly braces like functions
	if sum&1 == 0 {
		fmt.Printf("%v is even\n", sum)
	} else {
		fmt.Printf("%v is odd\n", sum)
	}

	// Shorthand If
	if Truth := sum % 5; Truth < 3 {
	} else {
	}
	// unlike regular if 'Truth' variable will not be available
	// out of this scope like for loop 1st example

	// like C switch but Go switch doesn't have paranthesis and
	// doesn't require 'break'

	switch isEven := sum % 2; isEven {
	// also 'isEven' can't be accessed out of the scope
	case 0:
		fmt.Printf("%v is even\n", sum)
	default:
		fmt.Printf("%v is odd\n", sum)
	}

	// we can also write switch statement where 'isEven' can be accessed
	isEven := sum % 2
	switch isEven {
	}

	// switch without condition
	switch {
	case sum%2 == 0:
		fmt.Printf("%v is even\n", sum)
	default:
		fmt.Printf("%v is odd\n", sum)
	}

	// New concept: defer
	// Executes after surrounding function ruturn
	// but defer statements are evaluted immediately ignoring its position
	defer fmt.Println("Defer Statement First")
	fmt.Println("Normal Statement Second")

	// defer are pushed to stack and executes in LIFO orders
	defer fmt.Print("\n")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v ", i)
	}
}

// func FUNCTION_NAME (arg1 ARGTYPE, arg2 ARGTYPE) RETURNTYPE { }
// NB: Opening Curly Braces '{' has to start in the same line of the function declaration
// Function 01 Example
func Div(x int, y int) (int, float32) {
	z := x / y                    // int [/*+-] int = int
	zf := float32(x) / float32(y) // float [/*+-] float = float
	return z, zf
}

// Naked Return Type where returned value is choosed by Variable name
// Function 02 Example
func ToFloat(x int) (xf float32) { //  we have set 'xf' as return value
	xf = float32(x) // rather than creating new variable we assign the value the returned variable
	return          // we can return new variable or just nil for 'xf'
}
