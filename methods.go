package main

import "fmt"

type Person struct {
	name string
	age int
}

// func (<receiver_name> <receiver_type>) <func_name>(args) <return_type>
func (p *Person) introduce() string {
	return "hello, my name is " + p.name
}

func main() {
	p := Person{"foo", 25}

	fmt.Println(p.introduce())
}
