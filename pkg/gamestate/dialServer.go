package gamestate

import (
	"log"
	"net"

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
		switch mrpItem.GetRequest() { /*
			case "TILE":
				json.Unmarshal([]byte(mrpItem.GetBody()), &myTileMap)
				UpdateGamestateFromServer()

			case "UNIT":
				var unitAdded units.IsUnit
				var myUnit *units.Unit
				err := json.Unmarshal([]byte(mrpItem.GetBody()), &myUnit)
				if err != nil {
					log.Fatal(err)
				}

				switch myUnit.GetType() {
				case "destroyer":
					unitAdded = units.NewDestroyer(myUnit, "")
				default:
					unitAdded = myUnit
				}
				myUnitMapTemp = append(myUnitMapTemp, unitAdded)

			case "END":
				PushUnitMap()

				UpdateGamestateFromServer()*/
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
