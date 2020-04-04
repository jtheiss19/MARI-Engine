package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
)

func NewBullet(conn net.Conn) *elements.Element {
	bullet := &elements.Element{}

	bullet.XPos = 0
	bullet.YPos = 0

	bullet.Active = true

	bullet.UniqueName = "player"

	//--FIRST ORDER--------------------------------------------//

	aPos := advancePos.NewAdvancePosition(bullet)
	bullet.AddComponent(aPos)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(bullet, "carrier.png")
	bullet.AddComponent(sr)

	replic := playerControl.NewReplicator(bullet, conn)
	bullet.AddComponent(replic)

	coli := physics.NewCollider(bullet)
	bullet.AddComponent(coli)

	rot := render.NewRotator(bullet)
	bullet.AddComponent(rot)

	//--THIRD ORDER--------------------------------------------//

	return bullet
}
