package gamestate

import "github.com/jtheiss19/project-undying/pkg/units"

var myUnitMap []*units.Destroyer

func GetUnitMap() []*units.Destroyer {
	return myUnitMap
}

func SpawnUnit() {
	myDestroyer := units.NewDestroyer()
	myUnitMap = append(myUnitMap, myDestroyer)
}
