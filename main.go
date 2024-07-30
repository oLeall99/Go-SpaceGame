package main

import (
	"go-game/game"
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := game.NewGame()

	iconFile, error := os.Open("assets/player.png")
	if error != nil {
		panic(error)
	}
	defer iconFile.Close()
	iconImage, error := png.Decode(iconFile)
	if error != nil {
		panic(error)
	}

	// Define o ícone da janela
	ebiten.SetWindowIcon([]image.Image{iconImage})

	// Configurações da janela
	ebiten.SetWindowTitle(" Go Space Game")
	ebiten.SetFullscreen(false)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
