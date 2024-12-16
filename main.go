// Source: https://go.dev/tour
package main

// NB: import(...) for importing more than one packages
import (
	"fmt"
	"math"
)

// Const has to be declared by '=' not this':=', Can be local or global
// NB: const(...) for declaring more than one constants
const GLOBAL_CONST = 5

func main() {
	/////////////////
	// Chapter 01: Basics
	/////////////////

	// Variable can be strongly typed with 'var'
	// or auto typed with ':='
	var v1 int = 6
	var v2 int = 4
	i, j := Div(v1, v2)
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
	// string Iteration

	myString := "Arafat"
	for i, r := range myString {
		fmt.Printf("%c : %v\n", myString[i], r)
	}
	// Alternative:
	// for i := 0; i < len(myString); i++ {}

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
	// unlike regular 'if', 'Truth' variable will not be available
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

	// defer fmt.Println("Defer Statement First")
	fmt.Println("Normal Statement Second")

	// defer are pushed to stack and executes in LIFO orders

	/// Commented out for ease of console logs
	// defer fmt.Print("\n")
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%v ", i)
	// }

	////////////////////////////////
	// Chapter 03: struct, slice & maps
	////////////////////////////////

	// Very much like C pointers, but doesn't have pointer arithmetic
	var p *int = &v1
	// or
	// p := &v1

	// 'p' holds address of 'v1' and can be accessed with '*'
	fmt.Printf("Addr: %p Value: %v\n", p, *p)
	// Changing value of '*p' will change value of 'v1' as they share same addrress
	*p = 4
	fmt.Printf("Addr: %p Value: %v Ref Value: %v\n", p, *p, v1)

	// To create struct object we can use 'var' or ':='
	myObj := myStruct{name: "test", id: 69}
	fmt.Println(myObj.name, myObj.id)

	// Static Arrays, has to be fixed size and can't be resized
	// Syntax: Variable := [Size] Type { elements }
	// Size can be ommited if the number elements are fixed
	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19}
	// var primes [5]int

	// We can slice array like Js or Python
	// NB: if slice range is 5:10 then array elements will be [5, 10)
	fmt.Println("Last 5 primes: ", primes[5:10])

	// array slice are referenced by shallow copy not a deep copy
	// so any changes made to slice array will affect the main array
	last3Primes := primes[7:10]
	fmt.Println("Last 3 primes: ", last3Primes)
	last3Primes[1] = 23
	last3Primes[2] = 29
	// last3Primes[1], last3Primes[2] = 23, 29 // Also valid oneliner
	fmt.Println("\nAfter changing the slice \nLast 5 primes: ", primes[5:10])
	fmt.Println("Last 3 primes: ", last3Primes)

	// 'len' for length and 'cap' for capacity
	// 'len()' is to get the number of elements
	// 'cap()' is to get the max number of elements can it hold
	// Dynamic Array can be made with 'make()' function
	// Syntax: make(TYPE, SIZE, CAPACITY) // CAPACITY is optional
	dynArr := make([]int, 2, 3)
	dynArr[0], dynArr[1] = 2, 3
	fmt.Printf("Size: %d Cap: %d Array %v\n", len(dynArr), cap(dynArr), dynArr)
	// Make sure to reasisgn the array after append
	dynArr = append(dynArr, 4)
	dynArr = append(dynArr, 5)
	// After append, capacity of dynamic array will double
	fmt.Printf("Size: %d Cap: %d Array %v\n", len(dynArr), cap(dynArr), dynArr)

	// 'range' kw acts like both Js 'in' & 'of' kws together
	// works with array, map
	// i(Index), v(Value)
	for i, v := range dynArr {
		fmt.Printf("(%v,%v) ", i, v)
	}
	println()

	// 'map' is key value pair data structure
	myMap := map[string]int{ // Static Map Literal way for small and known size
		"00": 0, // have to populate to work without that, in execution it will panic
	}
	// Value aissgnment and retrival are same as Js map
	myMap["01"] = 1
	myMap["10"] = 2
	myMap["11"] = 3

	fmt.Printf("len: %v Map: %v\n", len(myMap), myMap)

	for k, v := range myMap {
		fmt.Printf("0b%v = %v\n", k, v)
	}
	println()
	// Deleting Key
	delete(myMap, "00")

	// Test if value exisits, 'ok' is Ture if key exists in the map
	elem, ok := myMap["00"]
	if !ok {
		fmt.Printf("key [00] doesn't exist in the map, elem: %v\n", elem)
	}

	///////////////////////////
	// Methods and Interfaces//
	///////////////////////////
	mainStruct := myStruct{
		name: "Safin",
		id:   69,
	}
	// Passing as pointer
	revStruct := CreateRevStruct(&mainStruct)
	fmt.Printf("Main Addr: %p, Returned Addr: %p \n", &mainStruct, revStruct)

	// Using Struct specified Method
	conStruct := mainStruct.concate(revStruct)
	fmt.Printf("Concatenated Struct: %v\n", conStruct)

	// We can set interface type and assign corresponding struct to it
	// This Type of variables can only access methods only defined in the interface
	var Person1 myType = &myStruct{"Safin", 22}
	var Person2 myType = &myStruct{"Safin", 25}

	if Person1.insertUser() {
		fmt.Printf("User Inserted: %v\n", Person1)
		// can't access 'Person1.name' as it is not in interface
	}
	if Person2.insertUser() {
		fmt.Printf("User Inserted: %v\n", Person2)
	} else {
		fmt.Printf("User Already Exists: %v\n", Person2)
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
	xf = float32(x) // rather than creating new variable we assign the value to returned variable
	return          // we can return new variable or just nil for 'xf'
}

// For creating struct we have to use 'type' kw. More on that later
type myStruct struct {
	name string
	id   uint32
}

// Taking Pointer as Arg and Returning struct as pointer
func CreateRevStruct(V *myStruct) *myStruct {
	revName := ""
	for _, r := range V.name { // '_' used underscore to avoid 'Unused Variable' warning
		revName = string(r) + revName
	}

	var revID uint32 = 0
	n := V.id
	for n != 0 {
		digit := n % 10
		revID = revID*10 + digit
		n /= 10
	}
	funcStruct := myStruct{name: revName, id: revID}
	fmt.Printf("Arg Addr: %p, Return Addr: %p\n", V, &funcStruct)
	return &funcStruct
}

// Method can be defined for struct
func (M *myStruct) concate(m *myStruct) myStruct {
	fmt.Printf("Method Addr: %p\n", M)
	Name := M.name + m.name
	i := int(math.Log10(float64(m.id))) + 1
	ID := M.id*uint32(math.Pow(10, float64(i))) + m.id
	return myStruct{name: Name, id: ID}
}

// Skeleton of a Structure and objects of this types can only methods below
type myType interface {
	concate(*myStruct) myStruct
	insertUser() bool
}

// Stores 'name' of myStruct Object
var USER []string

func (M *myStruct) insertUser() bool {
	name := M.name
	for _, s := range USER {
		if s == name {
			return false
		}
	}
	USER = append(USER, name)
	return true
}
