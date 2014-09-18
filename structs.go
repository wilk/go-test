package main

import "fmt"

// type <struct_name> struct
type Person struct {
	name string
	age int
	gender string
}

func main() {
	// can be initialized sequentially
	foo := Person{"foo", 26, "male"}
	fmt.Println(foo)

	// otherwise in a dict form
	bar := Person{name: "bar", age: 46, gender: "female"}
	fmt.Println(bar)

	// internal properties called with a dot
	fmt.Println(foo.name, bar.name)

	// can be passed by reference
	fb := &foo
	// and go will dereference it for ya =)
	fmt.Println(fb.age)
}
