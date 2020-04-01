package gameloop

import (
	"fmt"

	"github.com/jtheiss19/project-undying/pkg/player"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

var PlrView *player.PlrView = player.NewPlrView()

func Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	controls(gamestate.GetServerConnection())

	tileCount := 0
	for _, myTile := range gamestate.GetTileMap() {
		x, y := myTile.GetPos()
		if PlrView.CanView(x, y) {
			x, y := PlrView.GetPos()
			myTile.Draw(screen, -x, -y)
			tileCount++
		}
	}

	for _, myUnit := range gamestate.GetUnitMap() {
		myUnit.Update()
		x, y := myUnit.GetPos()
		if PlrView.CanView(x, y) {
			x, y := PlrView.GetPos()
			myUnit.Draw(screen, -x, -y)
			tileCount++
		}
	}

	msg := fmt.Sprintf(" TPS: %0.2f \n FPS: %0.2f \n Number of Things Drawn: %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), tileCount)
	ebitenutil.DebugPrint(screen, msg)

	return nil
}
