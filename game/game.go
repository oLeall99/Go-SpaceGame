package game

import (
	"fmt"
	"go-game/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	stars            []*Star
	meteorSpawnTimer *Timer
	starSpawnTimer   *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starSpawnTimer:   NewTimer(14),
	}
	player := NewPlayer(g)
	g.player = player
	g.score = 0
	return g
}

// Atualiza a lógica do jogo
// A função update é chamada 60 vezes por segundo pela lib ebiten
func (g *Game) Update() error {

	g.player.Update() // Atualiza o player

	// Atualiza os lasers
	for _, l := range g.lasers {
		l.Update()
	}

	// Verifica se o timer de spawn de estrelas já acabou
	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		s := NewStar()
		g.stars = append(g.stars, s)
		g.starSpawnTimer.Reset()
	}

	// Atualização das estrelas
	for _, s := range g.stars {
		s.Update()
	}

	// Verifica se o timer de spawn de meteoros já acabou
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
		g.meteorSpawnTimer.Reset()
	}

	// Atualização dos meteoros
	for _, m := range g.meteors {
		m.Update()
	}

	// Verifica a colisão do player com os meteoros
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			fmt.Println("você perdeu")
			g.Reset()
		}
	}

	// Verifica a colisão dos meteoros com lasers
	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}
	}

	return nil
}

// Desenha o objetos na tela
// A função update é chamada 60 vezes por segundo pela lib ebiten
func (g *Game) Draw(screen *ebiten.Image) {
	for _, s := range g.stars {
		s.Draw(screen)
	}

	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), assets.FontUi, 20, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Adiciona laser na lista de laser do jogo
func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
