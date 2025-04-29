package main

import (
	"fmt"
)

func main() {
	// Boolean type
	var isGoFun bool = true
	fmt.Println("Boolean:", isGoFun)

	// Integer types
	var age int = 30                // Default int (platform-dependent size)
	var smallInt int8 = 120         // 8-bit integer
	var largeInt int64 = 9876543210 // 64-bit integer
	fmt.Println("Integers:", age, smallInt, largeInt)

	// Unsigned integers
	var u uint = 100       // unsigned int (non-negative)
	var u8 uint8 = 255     // 8-bit unsigned integer
	var u16 uint16 = 65535 // 16-bit unsigned integer
	fmt.Println("Unsigned Integers:", u, u8, u16)

	// Float types
	var price float32 = 19.99
	var pi float64 = 3.14159265359
	fmt.Println("Floats:", price, pi)

	// Complex numbers
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 5.5 + 6.6i
	fmt.Println("Complex numbers:", c1, c2)

	// String type
	var name string = "GoLang"
	fmt.Println("String:", name)

	// Rune (represents a Unicode code point)
	var letter rune = 'G'
	fmt.Println("Rune:", letter)

	// Byte (alias for uint8)
	var b byte = 'A'
	fmt.Println("Byte:", b)

	// Type inference
	inferred := "I am inferred as a string"
	fmt.Println("Inferred Type:", inferred)

	// Constants
	const piConst = 3.14
	fmt.Println("Constant:", piConst)
}
