package sample

import (
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
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
