package server

import (
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

func init() {
	gamestate.NewMap()
}

//Server starts a server on the selected port and acts
//as the main entrance into the server package.
func Server(port string) {

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	fmt.Println("Server Now Listening on Port:", port)

	newConnSignal := make(chan string)

	for {
		go session(server, newConnSignal)
		fmt.Println(<-newConnSignal)
	}
}

func session(ln net.Listener, newConnSignal chan string) {
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal <- "New Connection Made"

	closeConnection := make(chan string)

	go mrp.ReadMRPFromConn(conn, handleMRP)

	fmt.Println(<-closeConnection)
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {

}
