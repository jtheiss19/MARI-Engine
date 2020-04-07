package main

import (
	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 1
	addr         = "localhost:8080" //set to "" to launch game in single player
)

func main() {

	myScreen := objects.NewScreen(-screenWidth/2, -screenHeight/2)
	gamestate.SetScreen(myScreen)

	gamestate.SetWindowSize(screenWidth, screenHeight)
	gamestate.SetWindowScale(screenScale)

	if addr == "" {
		gamemap.NewWorld()
		newPlayer := objects.NewPlayer(nil)
		newPlayer.ID = "0"
		newPlayer.UniqueName = newPlayer.UniqueName + newPlayer.ID
		gamestate.AddElemToChunk(newPlayer, 0, 0)
	} else {
		gamestate.SetMultiplayerAddress(addr)
	}

	gamestate.StartClient()
}
