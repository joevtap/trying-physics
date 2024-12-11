package main

import (
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joevtap/trying-physics/ball"
	"github.com/joevtap/trying-physics/body"
	"github.com/joevtap/trying-physics/constants"
	"github.com/joevtap/trying-physics/math"
	"github.com/joevtap/trying-physics/platform"
	"github.com/joevtap/trying-physics/player"
)

type Game struct {
	balls    []ball.Ball
	fan      platform.FanPlatform
	platform platform.Platform

	resetPressed bool

	player player.Player
}

func main() {
	game := &Game{}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Trying Physics")

	game.Init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Init() {
	for i := range 4 {
		random := rand.Float64() * 100

		radius := random / 4

		if radius < 5 {
			radius = 5
		}

		mass := random

		if mass < 5 {
			mass = 5
		}

		g.balls = append(g.balls, ball.Ball{
			Body: body.DynamicBody{
				Position:     math.Vector2D{X: float64(100 * (i + 1)), Y: 10},
				Velocity:     math.Vector2D{X: 0, Y: 0},
				Acceleration: math.Vector2D{X: 0, Y: 0},
				Mass:         mass,
			},
			Radius: radius,
		})

		g.fan = platform.FanPlatform{
			Body: body.StaticBody{
				Position: math.Vector2D{X: constants.ScreenWidth/2 + 70, Y: 0 + 250},
				Width:    50,
				Height:   constants.ScreenHeight - 250,
			},
		}

		g.platform = platform.Platform{
			Body: body.StaticBody{
				Position: math.Vector2D{X: 50, Y: 200},
				Width:    200,
				Height:   20,
			},
		}

		g.player = player.Player{
			Body: body.DynamicBody{
				Position:     math.Vector2D{X: 50, Y: constants.ScreenHeight - 100},
				Velocity:     math.Vector2D{X: 0, Y: 0},
				Acceleration: math.Vector2D{X: 0, Y: 0},
				Mass:         1,
			},
			Width:  20,
			Height: 50,
		}
	}
}

func (g *Game) Reset() {
	g.balls = nil
	g.Init()
}

func (g *Game) Update() error {

	for i := range g.balls {
		g.balls[i].Update()
		g.balls[i].Body.ApplyGravity(constants.Gravity)

		g.platform.CollideWithBall(&g.balls[i])
		g.fan.ApplyWindToBall(&g.balls[i])

	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		if !g.resetPressed {
			g.resetPressed = true
			g.Reset()
		}
	} else {
		g.resetPressed = false
	}

	g.player.Update()
	g.player.Body.ApplyGravity(constants.Gravity)
	g.platform.CollideWithPlayer(&g.player)
	g.fan.ApplyWindToPlayer(&g.player)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.balls {
		g.balls[i].Draw(screen)
	}

	g.fan.Draw(screen)
	g.platform.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
