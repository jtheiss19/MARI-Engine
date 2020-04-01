package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/gameloop"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 2
)

func main() {
	if err := ebiten.Run(gameloop.Update, screenWidth/screenScale, screenHeight/screenScale, screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
