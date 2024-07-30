package game

import (
	"go-game/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]
	speed := 5.0

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -5,
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y da imagem na tela
	op.GeoM.Translate(s.position.X, s.position.Y)

	// Desenha a imagem na tela
	screen.DrawImage(s.image, op)
}
