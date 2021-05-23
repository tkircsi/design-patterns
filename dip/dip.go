package dip

import "fmt"

func Main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate("John")
}

// Dependency Inversion Principle
// HLM(High Level Module) should not depend on LLM(Low Level Module)
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

// Low-Level Module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations,
		Info{parent, Parent, child},
		Info{child, Child, parent})
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// High-Level Module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate(name string) {
	// break DIP
	// relations := r.relationships.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called ", rel.to.name)
	// 	}
	// }
	childs := r.browser.FindAllChildrenOf(name)
	for _, child := range childs {
		fmt.Println("John has a child called ", child.name)
	}
}
