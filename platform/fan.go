package platform

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/joevtap/trying-physics/ball"
	"github.com/joevtap/trying-physics/body"
	"github.com/joevtap/trying-physics/constants"
	"github.com/joevtap/trying-physics/math"
	"github.com/joevtap/trying-physics/player"
)

type FanPlatform struct {
	Body body.StaticBody
}

func (p *FanPlatform) Update() {}

func (p *FanPlatform) ApplyWindToBall(ball *ball.Ball) {
	if ball.Body.Position.X+ball.Radius > p.Body.Position.X && ball.Body.Position.X-ball.Radius < p.Body.Position.X+p.Body.Width {
		if ball.Body.Position.Y+ball.Radius > p.Body.Position.Y-p.Body.Height && ball.Body.Position.Y-ball.Radius < p.Body.Position.Y+p.Body.Height {
			windForce := math.Vector2D{X: 0, Y: -constants.Gravity * ball.Body.Mass * 1.2}
			ball.Body.ApplyForce(windForce)
		}
	}
}

func (p *FanPlatform) ApplyWindToPlayer(player *player.Player) {

	if player.Body.Position.X+player.Width > p.Body.Position.X && player.Body.Position.X < p.Body.Position.X+p.Body.Width {
		if player.Body.Position.Y+player.Height > p.Body.Position.Y-p.Body.Height && player.Body.Position.Y < p.Body.Position.Y+p.Body.Height {
			windForce := math.Vector2D{X: 0, Y: -constants.Gravity * player.Body.Mass * 1.2}
			player.Body.ApplyForce(windForce)

			player.Body.Velocity.X *= 0.5
		}
	}
}

func (p *FanPlatform) Draw(dst *ebiten.Image) {
	vector.StrokeRect(dst, float32(p.Body.Position.X), float32(p.Body.Position.Y), float32(p.Body.Width), float32(p.Body.Height), 5, color.RGBA{0, 255, 0, 255}, false)
}
