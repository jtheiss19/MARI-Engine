package thirdorder

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/gamestate"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/firstorder"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/secondorder"
)

type Explosion struct {
	container *elements.Element
	Type      string
	coliData  elements.Component
	damData   elements.Component
}

func init() {
	var comp = new(Explosion)
	gamestate.MRPMAP["Explosion"] = comp
}

func (explo *Explosion) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewExplosion(finalElem)
	finalElem.AddComponent(myComp)
}

func NewExplosion(container *elements.Element) *Explosion {
	return &Explosion{
		container: container,
		Type:      "Explosion",
		coliData:  container.GetComponent(new(secondorder.Collider)),
		damData:   container.GetComponent(new(firstorder.Damage)),
	}
}

func (explo *Explosion) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

func (explo *Explosion) OnMerge(compM elements.Component) error {
	return nil
}

func (explo *Explosion) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (explo *Explosion) OnCheck(elemC *elements.Element) error {
	return nil
}

func (explo *Explosion) OnUpdateServer() error {
	if explo.coliData.(*secondorder.Collider).HasCollided {
		for _, elem := range explo.coliData.(*secondorder.Collider).GetObjectsHit() {
			elemHP := elem.GetComponent(new(firstorder.Health))
			if elemHP != nil && explo.damData != nil {
				elemHP.(*firstorder.Health).TakeDamage(explo.damData.(*firstorder.Damage).Attack)
			}
		}

		gamestate.RemoveElem(explo.container)
	}

	if explo.container.GetComponent(new(secondorder.MoveTo)) == nil {
		gamestate.RemoveElem(explo.container)
	}
	return nil
}

func (explo *Explosion) SetContainer(container *elements.Element) error {
	explo.container = container
	explo.coliData = container.GetComponent(new(secondorder.Collider))
	explo.damData = container.GetComponent(new(firstorder.Damage))
	return nil
}

func (explo *Explosion) MakeCopy() elements.Component {
	myComp := *explo
	return &myComp
}
