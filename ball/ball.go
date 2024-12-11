package ball

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/joevtap/trying-physics/body"
	"github.com/joevtap/trying-physics/constants"
)

type Ball struct {
	Body   body.DynamicBody
	Radius float64
}

func (b *Ball) Update() {
	b.Body.Update()

	if b.Body.Position.Y+b.Radius > constants.ScreenHeight {
		b.Body.Velocity.Y *= -0.5
		b.Body.Position.Y = constants.ScreenHeight - b.Radius
	}

	if b.Body.Position.X+b.Radius > constants.ScreenWidth {
		b.Body.Position.X = 0 + b.Radius
	}
}

func (b *Ball) Draw(dst *ebiten.Image) {
	b.Body.Draw(dst)

	vector.DrawFilledCircle(dst, float32(b.Body.Position.X), float32(b.Body.Position.Y), float32(b.Radius), color.RGBA{255, 0, 100, 255}, false)
}
