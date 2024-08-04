package main

import (
	"fmt"
)

// Factory Design Pattern
// - A component responsible solely for the wholesale (not piecewise) creation of objects

type Person interface {
	SayHello()
}

type person struct {
	name     string
	age      int
	eyeCount int
}

type tiredPerson struct {
	person
}

func (tp *tiredPerson) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d years old. I am tired\n", tp.name, tp.age)
}

func NewPerson(name string, age int) Person {
	if age > 60 {
		return &tiredPerson{person{name, age, 2}}
	}
	return &person{name, age, 2}
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", p.name, p.age)
}

func main() {
	p := NewPerson("John", 65)
	p.SayHello()
}
