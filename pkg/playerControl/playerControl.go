package playerControl

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/render"
)

type KeyboardMover struct {
	Container *elements.Element
	Speed     float64

	Sr *render.SpriteRenderer
}

func NewKeyboardMover(container *elements.Element, speed float64) *KeyboardMover {
	return &KeyboardMover{
		Container: container,
		Speed:     speed,
		Sr:        container.GetComponent(&render.SpriteRenderer{}).(*render.SpriteRenderer),
	}
}

func (mover *KeyboardMover) OnDraw(screen *ebiten.Image) error {
	return nil
}

func (mover *KeyboardMover) OnUpdate() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		//myMRP := mrp.NewMRP([]byte("SPAWN"), []byte("test"), []byte("test"))
		//conn.Write(myMRP.MRPToByte())
		//gamestate.SpawnUnit("0")
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		mover.Container.XPos -= mover.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		mover.Container.XPos += mover.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		mover.Container.YPos -= mover.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		mover.Container.YPos += mover.Speed
	}

	return nil
}
