package gameloop

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

func controls() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		gamestate.SpawnUnit()
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		PlrView.MovePlrViewBy(-1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		PlrView.MovePlrViewBy(0, -1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		PlrView.MovePlrViewBy(0, 1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		PlrView.MovePlrViewBy(1, 0)
	}
}
