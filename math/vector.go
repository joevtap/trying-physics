package math

import "math"

type Vector2D struct {
	X, Y float64
}

func (v0 Vector2D) Add(v1 Vector2D) Vector2D {
	return Vector2D{v0.X + v1.X, v0.Y + v1.Y}
}

func (v0 Vector2D) Sub(v1 Vector2D) Vector2D {
	return Vector2D{v0.X - v1.X, v0.Y - v1.Y}
}

func (v0 Vector2D) Mult(s float64) Vector2D {
	return Vector2D{v0.X * s, v0.Y * s}
}

func (v0 Vector2D) Div(s float64) Vector2D {
	return Vector2D{v0.X / s, v0.Y / s}
}

func (v0 Vector2D) Normalize() Vector2D {
	l := v0.Mag()
	return Vector2D{v0.X / l, v0.Y / l}
}

func (v0 Vector2D) Mag() float64 {
	return math.Sqrt(v0.X*v0.X + v0.Y*v0.Y)
}

func (v0 Vector2D) Dot(v1 Vector2D) float64 {
	return v0.X*v1.X + v0.Y*v1.Y
}
