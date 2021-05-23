package lsp

// Liskov Substitution Principle
// You should be able to substitute an embedding type in place of its embedded part

import "fmt"

func Main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)

	sq2 := NewSquare2(5)
	r := sq2.Rectangle()
	UseIt(&r)
}

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.height = size
	sq.width = size
	return &sq
}

// violates LSP
func (s *Square) SetWidth(size int) {
	s.width = size
	s.height = size
}

func (s *Square) SetHeight(size int) {
	s.width = size
	s.height = size
}

type Square2 struct {
	size int
}

func NewSquare2(size int) *Square2 {
	return &Square2{size}
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func UseIt(s Sized) {
	width := s.GetWidth()
	s.SetHeight(10)
	expectedArea := 10 * width
	actualArea := s.GetWidth() * s.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, " but we got ", actualArea, "\n")
}
