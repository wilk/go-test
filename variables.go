package main

import "fmt"

func main() {
	// var <var_name> <var_type> = <default_value>
	var a string = "hello world"
	fmt.Println(a)

	// multi declaration: var <var_name>, <var2_name> = <default_value>, <default2_value>
	var b, c = 10, 20
	fmt.Println(b,c)

	// implicit type definition (go will infer the type)
	var d = true
	fmt.Println(d)

	// var can be omitted into a function, bypassed by ':='
	e := "test"
	fmt.Println(e)
}
