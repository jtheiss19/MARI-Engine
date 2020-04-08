package sample

import (
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/gamestate"
)

func init() {
	gamestate.ObjectMap["myElement"] = NewElement()
}

func NewElement() *elements.Element {
	myElement := &elements.Element{}

	//--FIRST ORDER--------------------------------------------//

	//--SECOND ORDER-------------------------------------------//

	//--THIRD ORDER--------------------------------------------//

	//--NETWORKING---------------------------------------------//

	return myElement
}
