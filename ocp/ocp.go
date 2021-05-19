package ocp

import (
	"fmt"
)

// open for extension, closed for modification

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
	// some parameters
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if p.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if p.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// specification pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type ColorSpecification struct {
	color Color
}

func (s ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == s.color
}

type AndSpecification struct {
	first, second Specification
}

func (s AndSpecification) IsSatisfied(p *Product) bool {
	return s.first.IsSatisfied(p) && s.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

func Main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("LargeGreen products (new):\n")
	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{greenSpec, largeSpec}
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large green\n", v.name)
	}
}
