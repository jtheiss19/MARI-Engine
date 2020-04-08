package objects

import (
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondorder"
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
