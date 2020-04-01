package gamestate

import (
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/tiles"
)

var elementList []*elements.Element
var elementListTemp []*elements.Element

func NewWorld() {
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			myWater := tiles.NewWater(float64(x*64), float64(y*64))
			elementList = append(elementList, myWater)
		}
	}
}

func GetWorld() []*elements.Element {
	return elementList
}

func PushElemMap() {
	elementList = elementListTemp
	elementListTemp = []*elements.Element{}
}
