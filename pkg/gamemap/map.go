package gamemap

import (
	"math/rand"
	"strconv"

	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

//NewWorld inits the elementList with elements representing
//water and a single player element. This is essentially a
//test enviroment.
func NewWorld() {
	for x := -10; x < 10; x++ {
		for y := -10; y < 10; y++ {
			myWater := objects.NewWater(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y))
			gamestate.AddElemToChunk(myWater, 0, 0)
			if rand.Intn(20) <= 0 {
				myIsland := objects.NewIsland(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y)+" ")
				gamestate.AddElemToChunk(myIsland, 0, 1)
			}
		}
	}
}
