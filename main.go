package main

import (
	"log"

	"github.com/Div9851/gba-emu/gba"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(480, 320)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(gba.New()); err != nil {
		log.Fatal(err)
	}
}
