package main

import "fmt"

// interfaces are collections of methods
type Behaviour interface {
	dancing() string
	singing() string
}

type Person struct {
	name string
	skill string
}

type Developer struct {
	name string
	dance string
	song string
}

// first dancing method, used by Person passing through Behaviour
func (p Person) dancing() string {
	return p.name + " is dancing " + p.skill
}

// first singing method, used by Person passing through Behaviour
func (p Person) singing() string {
	return p.name + " is singing " + p.skill
}

// second dancing method, used by Developer passing through Behaviour
func (d Developer) dancing() string {
	return d.name + " is not dancing " + d.dance
}

// second singing method, used by Person passing through Behaviour
func (d Developer) singing() string {
	return d.name + " is not singing " + d.song
}

// this function accepts a Behaviour. go will identify those methods that are attached to the
// current struct (Behaviour is just an interface and thus it's possible to pass different struct
// to the same function)
func activity(b Behaviour) {
	fmt.Println(b)
	fmt.Println(b.dancing())
	fmt.Println(b.singing())
}

func main() {
	p := Person{"wilk", "baciata"}
	activity(p)

	d := Developer{"wilk", "baciata", "singing in the rain"}
	activity(d)

	d2 := Developer{"wilk", "street dance", "riddiculus"}
	activity(d2)
}
