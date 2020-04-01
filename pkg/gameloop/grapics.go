package gameloop

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

func Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	controls()

	for _, myTile := range gamestate.GetTileMap() {
		myTile.Draw(screen)
	}

	for _, myUnit := range gamestate.GetUnitMap() {
		myUnit.Update()
		myUnit.Draw(screen)
	}

	msg := fmt.Sprintf("TPS: %0.2f \nFPS: %0.2f \nNumber of Ships: %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), len(gamestate.GetUnitMap()))
	ebitenutil.DebugPrint(screen, msg)

	return nil
}
