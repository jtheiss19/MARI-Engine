package main

import (
	"github.com/jtheiss19/project-undying/pkg/gamemap"
	"github.com/jtheiss19/project-undying/pkg/networking/server"
)

func main() {
	gamemap.NewWorld()

	server.StartServer("8080")
}
