package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
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

	shot := playerControl.NewShooter(player)
	player.AddComponent(shot)

	aPos := advancePos.NewAdvancePosition(player)
	player.AddComponent(aPos)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	coli := physics.NewCollider(player)
	player.AddComponent(coli)

	rot := render.NewRotator(player)
	player.AddComponent(rot)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	//--THIRD ORDER--------------------------------------------//

	return player
}
