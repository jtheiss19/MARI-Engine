package gamestate

import (
	"github.com/jtheiss19/project-undying/pkg/units"
)

var myUnitMap []units.IsUnit
var myUnitMapTemp []units.IsUnit

func GetUnitMap() []units.IsUnit {
	return myUnitMap
}

func SpawnUnit(owner string) {
	myDestroyer := units.NewDestroyer(nil, owner)
	myUnitMap = append(myUnitMap, myDestroyer)
}

func PushUnitMap() {
	myUnitMap = myUnitMapTemp
	myUnitMapTemp = []units.IsUnit{}
}
