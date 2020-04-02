package playerControl

import (
	"encoding/json"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

type Replicator struct {
	container *elements.Element
	conn      net.Conn
	Type      string
}

func NewReplicator(container *elements.Element, conn net.Conn) *Replicator {
	return &Replicator{
		container: container,
		conn:      conn,
		Type:      "Replicator",
	}
}

func (replic *Replicator) OnDraw(screen *ebiten.Image) error {
	return nil
}

func (replic *Replicator) OnUpdate() error {
	if replic.conn != nil {
		bytes, _ := json.Marshal(replic.container)
		myMRP := mrp.NewMRP([]byte("REPLIC"), []byte(bytes), []byte(replic.container.ID))
		replic.conn.Write(myMRP.MRPToByte())
	}
	return nil
}
