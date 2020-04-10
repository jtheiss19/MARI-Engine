package main

import (
	"github.com/jtheiss19/MARI-Engine/gamestate"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/objects"
	"github.com/jtheiss19/MARI-Engine/sampleGame/gamemap"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 0.75
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
		gamestate.AddElemToChunk(newPlayer, 0)
	} else {
		gamestate.SetMultiplayerAddress(addr)
	}

	gamestate.StartClient()
}
