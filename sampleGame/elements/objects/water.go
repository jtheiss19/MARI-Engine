package objects

import (
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/secondorder"
)

func NewWater(xpos float64, ypos float64, Name string) *elements.Element {
	water := &elements.Element{}

	water.XPos = xpos
	water.YPos = ypos

	water.Active = true

	water.UniqueName = Name

	sr := secondorder.NewSpriteRenderer(water, "water.png")
	water.AddComponent(sr)

	return water
}
