package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	var a string = "hello"
	fmt.Println(a)
	// Will throw out an error if the variable is not used
	var b, c int = 3, 6
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "shorthand" // The := syntax is shorthand for declaring and initializing a variable
	fmt.Println(f)

	const n float64 = 9000000000
	const m = 3e20 / n
	fmt.Println(int64(n))
	fmt.Println(math.Sin(m))
}