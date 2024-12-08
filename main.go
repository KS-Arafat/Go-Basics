// Source: https://go.dev/tour
package main

// NB: import(...) for importing more than one packages
import "fmt"

// Const has to be declared by '=' not this':=', Can be local or global
// NB: const(...) for declaring more than one constants
const GLOBAL_CONST = 5

func main() {
	i, j := Div(6, 4)
	// Like C printf with different types of flags
	fmt.Printf("Result %T : %v, %T : %v\n", i, i, j, j)
	a := ToFloat(6)
	fmt.Printf("Result %T : %v\n", a, a)

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
