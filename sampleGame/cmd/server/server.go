package main

import (
	"net"

	"github.com/jtheiss19/MARI-Engine/gamestate"
	"github.com/jtheiss19/MARI-Engine/networking/server"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/objects"
	"github.com/jtheiss19/MARI-Engine/sampleGame/gamemap"
)

func main() {
	gamemap.NewWorld()

	server.SetSpawnFunction(spawnStarterShip)

	server.StartServer("8080")
}

func spawnStarterShip(conn net.Conn, ID string) {
	newPlayer := objects.NewPlayer(conn)
	newPlayer.ID = ID
	newPlayer.UniqueName = newPlayer.UniqueName + ID
	gamestate.AddElemToChunk(newPlayer, 3)
}
