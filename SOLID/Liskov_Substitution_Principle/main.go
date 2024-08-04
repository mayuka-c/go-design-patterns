package main

import "fmt"

// LSP
// Liskov Substitution Principle
// Derived classes must be substitutable for their base classes
// If S is a subtype of T, then objects of type T may be replaced with objects of type S
// without altering any of the desirable properties of the program
// In other words, a derived class must extend the base class without changing its behavior
// In the example below, the Square struct violates the Liskov Substitution Principle
// because it changes the behavior of the Rectangle struct
// The SetWidth and SetHeight methods of the Square struct change the behavior of the Rectangle struct

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseIt(s Sized) {
	width := s.GetWidth()
	s.SetHeight(10)
	expectedArea := 10 * width
	actualArea := s.GetWidth() * s.GetHeight()
	fmt.Println("Expected area:", expectedArea, ". but got:", actualArea)
}

func main() {
	rc := Rectangle{2, 3}
	UseIt(&rc)

	sq := NewSquare(5)
	UseIt(sq)
}
