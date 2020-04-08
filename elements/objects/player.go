package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/firstorder"
	"github.com/jtheiss19/project-undying/pkg/elements/secondorder"

	"github.com/jtheiss19/project-undying/pkg/elements"
)

const (
	playerSpeed = 1
)

func NewPlayer(conn net.Conn) *elements.Element {
	player := &elements.Element{}

	player.XPos = 0
	player.YPos = 0

	player.Active = true

	player.UniqueName = "player"

	//--FIRST ORDER--------------------------------------------//

	aPos := firstorder.NewAdvancePosition(player, playerSpeed)
	player.AddComponent(aPos)

	hp := firstorder.NewHealth(player, 100)
	player.AddComponent(hp)

	//--SECOND ORDER-------------------------------------------//

	sr := secondorder.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	shot := secondorder.NewShooter(player)
	player.AddComponent(shot)

	mover := secondorder.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	coli := secondorder.NewCollider(player)
	player.AddComponent(coli)

	rot := secondorder.NewRotator(player)
	player.AddComponent(rot)

	//--THIRD ORDER--------------------------------------------//

	//--NETWORKING---------------------------------------------//

	replic := secondorder.NewReplicator(player, conn)
	player.AddComponent(replic)

	return player
}
