package main

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/server"
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
	gamestate.AddElemToChunk(newPlayer, 0, 3)
}
