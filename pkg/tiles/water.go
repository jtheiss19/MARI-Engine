package tiles

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/render"
)

var waterTexture *ebiten.Image

func NewWater(xpos float64, ypos float64) *elements.Element {
	water := &elements.Element{}

	water.XPos = xpos
	water.YPos = ypos

	water.Active = true

	sr := render.NewSpriteRenderer(water, "water.png", waterTexture)
	water.AddComponent(sr)

	waterTexture = sr.Tex

	return water
}
