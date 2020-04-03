package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/physics"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
)

const (
	playerSpeed = 1
)

func NewPlayer(conn net.Conn) *elements.Element {
	player := &elements.Element{}

	player.XPos = 1280 / 2
	player.YPos = 720 / 2

	player.Active = true

	player.UniqueName = "player"

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	coli := physics.NewCollider(player)
	player.AddComponent(coli)

	rot := render.NewRotator(player)
	player.AddComponent(rot)

	return player
}
