package server

import (
	"time"

	"github.com/jtheiss19/MARI-Engine/gamestate"
)

const tps = 60

func StartServer(port string) {
	go Server(port)

	var timeSinceLastUpdate int64
	var now int64
	count := 0
	for {

		time.Sleep((1000/tps - time.Duration(timeSinceLastUpdate)) * time.Millisecond)
		if count == 60 {
			//fmt.Println(timeSinceLastUpdate)
			count = 0
		}
		count++
		now = time.Now().UnixNano()

		world := gamestate.GetEntireWorld()

		for _, layer := range world {
			for _, elem := range layer {
				if elem != nil {
					if elem.Active {
						elem.UpdateServer()
					}
				}
			}
		}

		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
