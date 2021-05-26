package bridge

import "fmt"

func MainBridge() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	c := NewCircle(&raster, 10)
	c.Draw()
	c.Resize(2)
	c.Draw()

	c2 := NewCircle(&vector, 5)
	c2.Draw()
	c2.Resize(10)
	c2.Draw()
}

type Renderer interface {
	RenderCircle(radius int)
}

type VectorRenderer struct {
	// params for vector rendering
}

func (v *VectorRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing a vector circle with radius", radius)
}

type RasterRenderer struct {
	// params for raster rendering
}

func (r *RasterRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing a raster circle with radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   int
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor int) {
	c.radius *= factor
}

func NewCircle(renderer Renderer, radius int) *Circle {
	return &Circle{
		renderer: renderer,
		radius:   radius,
	}
}
