package render

import (
	"math"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

//Rotator is the component that handles all
//rendering of sprites onto the screen
type Rotator struct {
	container *elements.Element
	posData   elements.Component

	Type string
}

func init() {
	var rot = new(Rotator)
	gamestate.MRPMAP["Rotator"] = rot
}

//NewRotator creates a SpriteRenderer which
//is the component that handles all rendering of
//sprites onto the screen
func NewRotator(container *elements.Element) *Rotator {

	return &Rotator{
		container: container,
		posData:   container.GetComponent(new(advancePos.AdvancePosition)),
		Type:      "Rotator",
	}
}

func (rot *Rotator) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewRotator(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw Draws the stored texture file onto the screen
func (rot *Rotator) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate is used to qualify SpriteRenderer as a component
func (rot *Rotator) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (rot *Rotator) OnCheck(elemC *elements.Element) error {
	return nil
}

func (rot *Rotator) OnMerge(compM elements.Component) error {
	return nil
}

func (rot *Rotator) OnUpdateServer() error {

	if rot.container.YPos == rot.posData.(*advancePos.AdvancePosition).PrevY && rot.container.XPos == rot.posData.(*advancePos.AdvancePosition).PrevX {
	} else {
		rot.container.Rotation = math.Atan2((rot.container.YPos - rot.posData.(*advancePos.AdvancePosition).PrevY), (rot.container.XPos - rot.posData.(*advancePos.AdvancePosition).PrevX))
	}

	return nil
}
