package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"

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

	sendElemMap(conn)

	closeConnection := make(chan string)
	fmt.Println(<-closeConnection)
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "UNIT":
			sendElemMap(conn)
		}
	}
}

func sendElemMap(conn net.Conn) {
	myMap := gamestate.GetWorld()

	for _, myElem := range myMap {
		bytes, err := json.Marshal(myElem)
		if err != nil {
			log.Fatal(err)
		}

		myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
		conn.Write(myMRP.MRPToByte())
	}

	myMRP := mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
	conn.Write(myMRP.MRPToByte())
}
