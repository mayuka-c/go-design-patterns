package main

import (
	"fmt"
	"strings"

	builderfacet "github.com/mayuka-c/go-design-patterns/Creational-Design-Pattern/Builder/builder-facet"
	builderparameter "github.com/mayuka-c/go-design-patterns/Creational-Design-Pattern/Builder/builder-parameter"
)

// Builder Design Pattern
// - When piecewise object construction is complicated, provide an API for doing it succinctly
// - A builder is a separate component for building an object

const (
	indentSize = 2
)

type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName, HTMLElement{rootName, "", []HTMLElement{}}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(childName, childText string) {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	b := NewHTMLBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	c := NewHTMLBuilder("ul")
	c.AddChildFluent("li", "hello").AddChildFluent("li", "world")
	fmt.Println(c.String())

	// builder facet
	pb := builderfacet.NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(person)

	// build parameter
	email := builderparameter.SendEmail(func(b *builderparameter.EmailBuilder) {
		b.From("test@bar.com").To("test@foo.com").Subject("Hello").Body("Hello, this is a test email")
	})
	fmt.Println(email)
}
