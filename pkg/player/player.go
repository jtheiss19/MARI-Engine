package player

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/playerControl"
	"github.com/jtheiss19/project-undying/pkg/render"
)

const (
	playerSpeed = 1
)

var playerTexture *ebiten.Image

func NewPlayer() *elements.Element {
	player := &elements.Element{}

	player.XPos = 0
	player.YPos = 0

	player.Active = true

	sr := render.NewSpriteRenderer(player, "destroyer.png", playerTexture)
	player.AddComponent(sr)

	playerTexture = sr.Tex

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	return player
}
