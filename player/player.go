package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/joevtap/trying-physics/body"
	"github.com/joevtap/trying-physics/constants"
)

type Player struct {
	Body body.DynamicBody

	Width  float64
	Height float64
}

func (p *Player) Update() {
	p.Body.Update()

	if p.Body.Position.Y+p.Height > constants.ScreenHeight {
		p.Body.Velocity.Y *= -0.1
		p.Body.Position.Y = constants.ScreenHeight - p.Height

		p.Body.Velocity.X *= 0.2

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			p.Body.Velocity.Y = -20
		}

	}

	if p.Body.Position.X+p.Width > constants.ScreenWidth {
		p.Body.Position.X = 0 + p.Width
	}

	if p.Body.Position.X < 0 {
		p.Body.Position.X = constants.ScreenWidth - p.Width
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Body.Velocity.X = -5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Body.Velocity.X = 5
	}

}

func (p *Player) Draw(dst *ebiten.Image) {
	vector.DrawFilledRect(dst, float32(p.Body.Position.X), float32(p.Body.Position.Y), float32(p.Width), float32(p.Height), color.RGBA{255, 255, 0, 255}, false)
}
