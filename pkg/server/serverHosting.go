package server

import (
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/mrp"
)

var connectionList = make(map[int]net.Conn)

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

	connections := 0
	for {
		go session(server, newConnSignal, connections)
		fmt.Println(<-newConnSignal)
		connections++
	}
}

func session(ln net.Listener, newConnSignal chan string, sessionID int) {
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal <- "New Connection Made"

	connectionList[sessionID] = conn

	go mrp.ReadMRPFromConn(conn, handleMRP)

	sendTileMap(conn)

	closeConnection := make(chan string)
	fmt.Println(<-closeConnection)
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "UNIT":
			sendUnitMap(conn)
		}
	}
}

func sendTileMap(conn net.Conn) {
	//myMap := gamestate.GetTileMap()

	//bytes, _ := json.Marshal(myMap)

	//myMRP := mrp.NewMRP([]byte("TILE"), bytes, []byte("test"))

	//conn.Write(myMRP.MRPToByte())
}

func sendUnitMap(conn net.Conn) {
	//myMap := gamestate.GetUnitMap()

	//for _, myUnit := range myMap {
	//	bytes, _ := json.Marshal(myUnit)
	//
	//	myMRP := mrp.NewMRP([]byte("UNIT"), bytes, []byte("test"))

	//	conn.Write(myMRP.MRPToByte())
	//}

	//myMRP := mrp.NewMRP([]byte("END"), []byte("test"), []byte("test"))

	//conn.Write(myMRP.MRPToByte())
}
