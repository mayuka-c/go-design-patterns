package main

import "fmt"

// OCP
// Open for extension, closed for modification
// Once a module is written, it should be closed for modification
// But it should be open for extension

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for index, product := range products {
		if product.color == color {
			result = append(result, &products[index])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for index, product := range products {
		if product.size == size {
			result = append(result, &products[index])
		}
	}
	return result
}

// New way of doing things
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct {
}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for index := range products {
		if spec.IsSatisfied(&products[index]) {
			result = append(result, &products[index])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	f := Filter{}
	fmt.Println("Green products (old):")
	for _, product := range f.FilterByColor(products, green) {
		fmt.Printf("- %s is green\n", product.name)
	}

	// New way
	fmt.Println("Green products (new):")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, product := range bf.Filter(products, greenSpec) {
		fmt.Printf("- %s is green\n", product.name)
	}

	fmt.Println("Large products (new):")
	largeSpec := SizeSpecification{large}
	for _, product := range bf.Filter(products, largeSpec) {
		fmt.Printf("- %s is large\n", product.name)
	}

	fmt.Println("Large blue items (new):")
	largeBlueSpec := AndSpecification{largeSpec, ColorSpecification{blue}}
	for _, product := range bf.Filter(products, largeBlueSpec) {
		fmt.Printf("- %s is large and blue\n", product.name)
	}
}
