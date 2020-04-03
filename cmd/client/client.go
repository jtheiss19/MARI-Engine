package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/gameloop"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 1
)

func main() {

	gamestate.Dial("localhost:8080")

	time.Sleep(1 * time.Second)

	gameloop.MakeScreen()

	if err := ebiten.Run(gameloop.Update, screenWidth/screenScale, screenHeight/screenScale, screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
