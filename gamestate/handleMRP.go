package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/connection"
	"github.com/jtheiss19/MARI-Engine/networking/mrp"
)

var serverConnection net.Conn
var MRPMAP = make(map[string]elements.Component)

//Dial setsup a gamestate to be controlled by the server dialed
//via the address variable.
func Dial(address string) {
	var err error

	serverConnection, err = net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	go mrp.ReadMRPFromConn(serverConnection, HandleMRP)
}

func HandleMRP(mrpItem *mrp.MRP, conn net.Conn) {
	switch mrpItem.GetRequest() {
	case "ELEM":
		bytesMaster := []byte(mrpItem.GetBody())

		var finalElem = new(elements.Element)
		handleELEMCreates(bytesMaster, finalElem)

		if mrpItem.GetFooters()[0] == "NIL" {
			for _, layer := range GetEntireWorld() {
				for _, elem := range layer {
					if elem.UniqueName == finalElem.UniqueName {
						blacklistedNames = append(blacklistedNames, elem.UniqueName)
						RemoveElem(elem)
						break
					}
				}
			}
		} else {
			layerToAddOn, err := strconv.Atoi(mrpItem.GetFooters()[0])
			if err != nil {
				log.Fatal(err)
			}
			AddElemToChunk(finalElem, layerToAddOn)
		}

	case "REPLIC":
		world := GetEntireWorld()
		handleREPLIC(mrpItem, conn, world)

	case "ID":
		connection.SetID(mrpItem.GetBody())

	default:
		fmt.Println("Command Not Understood")
	}

}

func handleELEMCreates(bytesMaster []byte, finalElem *elements.Element) {

	var tempElem = *new(map[string]interface{})

	json.Unmarshal(bytesMaster, &tempElem)

	//fmt.Println(tempElem)

	test := tempElem["Components"].([]interface{})
	for _, comp := range test {

		if comp != nil {

			//var myComp elements.Component
			kindOfComp := comp.(map[string]interface{})["Type"].(string)
			myComp := MRPMAP[kindOfComp]
			if myComp != nil {
				myComp.MRP(finalElem, serverConnection)
			}
		}
	}

	json.Unmarshal(bytesMaster, &finalElem)
}

func handleREPLIC(mrpItem *mrp.MRP, conn net.Conn, world [][]*elements.Element) {
	for _, layer := range world {
		for _, elem := range layer {
			if elem.UniqueName == mrpItem.GetFooters()[0] {
				var elemTemp = new(elements.Element)
				handleELEMCreates([]byte(mrpItem.GetBody()), elemTemp)

				if elem.Check(elemTemp) == nil {
					elemTemp.Merge(elem)
				}

				break
			}
		}
	}

}
