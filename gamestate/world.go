package gamestate

import (
	"encoding/json"
	"log"
	"net"
	"strconv"

	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/mrp"
)

var connectionList = make(map[int]net.Conn)

func SendElem(conn net.Conn, elem *elements.Element, planeID int) {
	bytes, _ := json.Marshal(&elem)

	myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(strconv.Itoa(planeID)))
	conn.Write(myMRP.MRPToByte())
}

func NewConnection(conn net.Conn, ID int) {
	connectionList[ID] = conn
}

func UpdateElemToAll(elem *elements.Element, planeID int) {
	for _, client := range connectionList {
		SendElem(client, elem, planeID)
	}
}

func SendElemMap(conn net.Conn) {
	myMap := GetEntireWorld()

	for level, layer := range myMap {
		for _, elem := range layer {
			bytes, err := json.Marshal(elem)
			if err != nil {
				log.Fatal(err)
			}

			myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(strconv.Itoa(level)))
			conn.Write(myMRP.MRPToByte())

		}

	}
}
