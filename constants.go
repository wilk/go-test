package main

import "fmt"
import "math"

const a = 500000
const s = "hello world"

func main() {
	fmt.Println(s, a)

	// constants can be defined where a variable can
	const b = 2.10
	fmt.Println(b)

	// they don't have an explicit type
	const d = 3e20 / a
	fmt.Println(d)

	// until they are casted
	fmt.Println(int64(d)) // that's a type cast

	// in this case, Sin expect a float64 and 'a' is an integer
	// so, it casts to float64, moving the type of 'a' to its context
	fmt.Println(math.Sin(a))
}
