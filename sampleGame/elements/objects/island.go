package objects

import (
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/sampleGame/elements/secondorder"
)

func NewIsland(xpos float64, ypos float64, Name string) *elements.Element {
	Island := &elements.Element{}

	Island.XPos = xpos
	Island.YPos = ypos

	Island.Active = true

	Island.UniqueName = Name

	sr := secondorder.NewSpriteRenderer(Island, "island.png")
	Island.AddComponent(sr)

	coli := secondorder.NewCollider(Island)
	coli.Radius = 25
	Island.AddComponent(coli)

	return Island
}
