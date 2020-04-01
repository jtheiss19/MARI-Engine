package gameloop

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

func controls(conn net.Conn) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		myMRP := mrp.NewMRP([]byte("SPAWN"), []byte("test"), []byte("test"))
		conn.Write(myMRP.MRPToByte())
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
