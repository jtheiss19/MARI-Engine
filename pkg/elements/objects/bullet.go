package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
)

func NewBullet(conn net.Conn, DestX, DestY float64) *elements.Element {
	bullet := &elements.Element{}

	bullet.XPos = 0
	bullet.YPos = 0

	bullet.Active = true

	bullet.UniqueName = "MyBullet"

	//--FIRST ORDER--------------------------------------------//

	aPos := advancePos.NewAdvancePosition(bullet, 5)
	bullet.AddComponent(aPos)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(bullet, "carrier.png")
	bullet.AddComponent(sr)

	rot := render.NewRotator(bullet)
	bullet.AddComponent(rot)

	mov := playerControl.NewMoveTo(bullet, DestX, DestY)
	bullet.AddComponent(mov)

	coli := physics.NewCollider(bullet)
	bullet.AddComponent(coli)
	//--THIRD ORDER--------------------------------------------//

	replic := playerControl.NewReplicator(bullet, conn)
	bullet.AddComponent(replic)

	return bullet
}
