package objects

import (
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/connection"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/secondorder"
)

func NewScreen(xpos float64, ypos float64) *elements.Element {
	screen := &elements.Element{}

	screen.XPos = xpos
	screen.YPos = ypos

	screen.Active = true

	screen.UniqueName = "MySpecialScreen"
	screen.ID = connection.GetID()

	mover := secondorder.NewKeyboardMover(screen, 1)
	screen.AddComponent(mover)

	return screen
}
