package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/firstorder"
	"github.com/jtheiss19/project-undying/pkg/elements/secondorder"
	"github.com/jtheiss19/project-undying/pkg/elements/thirdorder"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

func init() {
	gamestate.ObjectMap["Bullet"] = NewBullet(nil, 0, 0)
}

func NewBullet(conn net.Conn, DestX, DestY float64) *elements.Element {
	bullet := &elements.Element{}

	bullet.XPos = 0
	bullet.YPos = 0

	bullet.Active = true

	bullet.UniqueName = "MyBullet"

	//--FIRST ORDER--------------------------------------------//

	aPos := firstorder.NewAdvancePosition(bullet, 5)
	bullet.AddComponent(aPos)

	dam := firstorder.NewDamage(bullet)
	bullet.AddComponent(dam)

	//--SECOND ORDER-------------------------------------------//

	sr := secondorder.NewSpriteRenderer(bullet, "carrier.png")
	bullet.AddComponent(sr)

	rot := secondorder.NewRotator(bullet)
	bullet.AddComponent(rot)

	coli := secondorder.NewCollider(bullet)
	bullet.AddComponent(coli)

	mov := secondorder.NewMoveTo(bullet, -400, -400)
	bullet.AddComponent(mov)

	//--THIRD ORDER--------------------------------------------//

	explo := thirdorder.NewExplosion(bullet)
	bullet.AddComponent(explo)

	replic := secondorder.NewReplicator(bullet, conn)
	bullet.AddComponent(replic)

	return bullet
}
