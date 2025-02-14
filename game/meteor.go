package game

import (
	"go-game/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))] // Pega uma imagem randomica da lista de meteoros
	speed := (rand.Float64() * 13)

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}
	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y da imagem na tela
	op.GeoM.Translate(m.position.X, m.position.Y)

	// Desenha a imagem na tela
	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.image.Bounds()

	return NewRect(m.position.X, m.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
