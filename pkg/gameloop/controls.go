package gameloop

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

func controls() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		gamestate.SpawnUnit()
	}
}
