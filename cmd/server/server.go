package main

import (
	"time"

	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/server"
)

const tps = 60

func main() {
	go server.Server("8080")

	var timeSinceLastUpdate int64
	for {
		time.Sleep((1000/tps - time.Duration(timeSinceLastUpdate)) * time.Millisecond)
		now := time.Now().UnixNano()
		for _, myUnit := range gamestate.GetUnitMap() {
			myUnit.Update()
		}
		timeSinceLastUpdate = (time.Now().UnixNano() - now) / 1000000
	}
}
