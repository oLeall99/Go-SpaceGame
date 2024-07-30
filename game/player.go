package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image        *ebiten.Image
	position     Vector
	game         *Game
	loadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	//halfH := float64(bounds.Dy())

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 600,
	}

	return &Player{
		image:        image,
		game:         game,
		position:     position,
		loadingTimer: NewTimer(15),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.position.X += speed
	}

	p.loadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.loadingTimer.IsReady() {
		p.loadingTimer.Reset()

		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)

	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posilção X e Y da imagem na tela
	op.GeoM.Translate(p.position.X, p.position.Y)

	//Desenha imagem na tela
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(p.position.X, p.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
