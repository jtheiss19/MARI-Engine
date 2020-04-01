package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/tiles"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

var serverConnection net.Conn

func Dial(address string) {
	var err error

	serverConnection, err = net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	go mrp.ReadMRPFromConn(serverConnection, handleMRP)
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "ELEM":
			var tempElem elements.Element
			var finalElem *elements.Element
			json.Unmarshal([]byte(mrpItem.GetBody()), &tempElem)

			switch tempElem.Type {
			case "water":
				finalElem = tiles.NewWater(tempElem.XPos, tempElem.YPos)
			default:
				fmt.Println("No Match Found for Tile Data Type:", tempElem.Type)
			}
			elementListTemp = append(elementListTemp, finalElem)
		case "END":
			PushElemMap()

			UpdateGamestateFromServer()
		}
	}
}

func UpdateGamestateFromServer() {
	myMRP := mrp.NewMRP([]byte("UNIT"), []byte("test"), []byte("test"))
	serverConnection.Write(myMRP.MRPToByte())
}

func GetServerConnection() net.Conn {
	return serverConnection
}
