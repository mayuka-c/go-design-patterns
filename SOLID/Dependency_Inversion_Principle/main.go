package main

import "fmt"

// DIP
// Dependency Inversion Principle
// High-level modules should not depend on low-level modules
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-level module - Data storage in DB
// it breaks if relations instead of slice, gets from DB
type Relationships struct {
	relations []Info
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

// High-level module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationShipBrowser
}

type RelationShipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.from.name == name && v.relationship == Parent {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Research) Investigate() {
	childrens := r.browser.FindAllChildrenOf("John")
	for _, rel := range childrens {
		fmt.Println("John has a child called", rel.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate()
}
