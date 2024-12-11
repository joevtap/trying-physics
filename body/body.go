package body

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joevtap/trying-physics/math"
)

type DynamicBody struct {
	Position     math.Vector2D
	Velocity     math.Vector2D
	Acceleration math.Vector2D
	Mass         float64
}

func (b *DynamicBody) Update() {
	b.Velocity = b.Velocity.Add(b.Acceleration)
	b.Position = b.Position.Add(b.Velocity)

	b.Acceleration = b.Acceleration.Mult(0)
}

func (b *DynamicBody) ApplyGravity(gravity float64) {
	g := math.Vector2D{X: 0, Y: gravity}
	b.ApplyForce(g.Mult(b.Mass))
}

func (b *DynamicBody) ApplyForce(force math.Vector2D) {
	b.Acceleration = b.Acceleration.Add(force.Div(b.Mass))
}

func (b DynamicBody) Draw(dst *ebiten.Image) {}

type StaticBody struct {
	Position math.Vector2D
	Width    float64
	Height   float64
}

func (b StaticBody) Draw(dst *ebiten.Image) {}
