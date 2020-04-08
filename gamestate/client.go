package gamestate

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/connection"
)

var (
	windowWidth  = 1280
	windowHeight = 720
	screenScale  = 1.0
	addr         = ""
)

var isServer = false

var myScreen *elements.Element

func SetScreen(screen *elements.Element) {
	myScreen = screen
}

func SetWindowSize(XPos int, YPos int) {
	windowWidth = XPos
	windowHeight = YPos
}

func SetMultiplayerAddress(address string) {
	addr = address
}

func SetWindowScale(scale float64) {
	screenScale = scale
}

func StartClient() {

	if addr != "" {
		Dial(addr)
	} else {
		isServer = true
		connection.SetID("0")
	}

	SetScreen(myScreen)

	if err := ebiten.Run(Update, int(float64(windowWidth)/screenScale), int(float64(windowHeight)/screenScale), screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
