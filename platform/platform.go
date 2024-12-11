package platform

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/joevtap/trying-physics/ball"
	"github.com/joevtap/trying-physics/body"
	"github.com/joevtap/trying-physics/player"
)

type Platform struct {
	Body body.StaticBody
}

func (p *Platform) Update() {}

func (p *Platform) CollideWithBall(ball *ball.Ball) {
	if ball.Body.Position.X+ball.Radius > p.Body.Position.X && ball.Body.Position.X-ball.Radius < p.Body.Position.X+p.Body.Width {
		if ball.Body.Position.Y+ball.Radius > p.Body.Position.Y {
			ball.Body.Velocity.Y *= -0.9
			ball.Body.Position.Y = p.Body.Position.Y - ball.Radius

		}
	}
}

func (p *Platform) CollideWithPlayer(player *player.Player) {
	if player.Body.Position.X+player.Width > p.Body.Position.X && player.Body.Position.X < p.Body.Position.X+p.Body.Width {
		if player.Body.Position.Y+player.Height > p.Body.Position.Y && player.Body.Position.Y < p.Body.Position.Y+p.Body.Height {
			player.Body.Velocity.Y *= -0.8
			player.Body.Position.Y = p.Body.Position.Y - player.Height

		}
	}
}

func (p *Platform) Draw(dst *ebiten.Image) {
	vector.StrokeRect(dst, float32(p.Body.Position.X), float32(p.Body.Position.Y), float32(p.Body.Width), float32(p.Body.Height), 5, color.RGBA{0, 255, 255, 255}, false)
}
