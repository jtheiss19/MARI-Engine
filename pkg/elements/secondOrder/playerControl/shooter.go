package playerControl

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type Shooter struct {
	container *elements.Element
	Type      string
	HasFired  bool
	DestX     float64
	DestY     float64
}

func init() {
	var shoot = new(Shooter)
	gamestate.MRPMAP["Shooter"] = shoot
}

//NewReplicator creates a Replicator which is
//the component that handles all replication
//of an element onto a server.
func NewShooter(container *elements.Element) *Shooter {
	return &Shooter{
		container: container,
		Type:      "Shooter",
	}
}

func (shoot *Shooter) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewShooter(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (shoot *Shooter) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (shoot *Shooter) OnUpdate(xOffset float64, yOffset float64) error {
	if shoot.container.ID != connection.GetID() {
		return nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shoot.HasFired = true
		w, h := ebiten.CursorPosition()
		//myScreen := gameScreen.GetScreen()
		shoot.DestX = float64(w) - xOffset
		shoot.DestY = float64(h) - yOffset
	}
	return nil
}

func (shoot *Shooter) OnCheck(elemC *elements.Element) error {
	return nil
}

func (shoot *Shooter) OnMerge(compM elements.Component) error {
	compM.(*Shooter).HasFired = shoot.HasFired
	compM.(*Shooter).DestX = shoot.DestX
	compM.(*Shooter).DestY = shoot.DestY
	return nil
}

func (shoot *Shooter) OnUpdateServer() error {
	if shoot.HasFired {
		//shoot.container.AddComponentPostInit(NewMoveTo(shoot.container, shoot.DestX, shoot.DestY))
		gamestate.AddUnitToWorld(NewBullet(nil, shoot.DestX, shoot.DestY))
		shoot.HasFired = false
		gamestate.PushChunks()
	}

	return nil
}

func NewBullet(conn net.Conn, DestX, DestY float64) *elements.Element {
	bullet := &elements.Element{}

	bullet.XPos = 0
	bullet.YPos = 0

	bullet.Active = true

	bullet.UniqueName = "MyBullet"

	//--FIRST ORDER--------------------------------------------//

	aPos := advancePos.NewAdvancePosition(bullet, 10)
	bullet.AddComponent(aPos)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(bullet, "carrier.png")
	bullet.AddComponent(sr)

	coli := physics.NewCollider(bullet)
	bullet.AddComponent(coli)

	rot := render.NewRotator(bullet)
	bullet.AddComponent(rot)

	mov := NewMoveTo(bullet, DestX, DestY)
	bullet.AddComponent(mov)

	//--THIRD ORDER--------------------------------------------//

	replic := NewReplicator(bullet, conn)
	bullet.AddComponent(replic)

	return bullet
}
