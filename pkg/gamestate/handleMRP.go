package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
	"github.com/jtheiss19/project-undying/pkg/networking/mrp"
)

var serverConnection net.Conn

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

func HandleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "ELEM":
			bytesMaster := []byte(mrpItem.GetBody())

			var finalElem = new(elements.Element)
			handleELEMCreates(bytesMaster, finalElem)

			elementListTemp = append(elementListTemp, finalElem)
			if len(elementListTemp) > 10 {
				PushElemMap()
			}

		case "REPLIC":
			for _, elem := range GetWorld() {
				if elem.ID == mrpItem.GetFooters()[0] {
					var elemTemp = new(elements.Element)
					handleELEMCreates([]byte(mrpItem.GetBody()), elemTemp)

					if elem.Check(elemTemp) == nil {
						handleELEMCreates([]byte(mrpItem.GetBody()), elem)
					}
					go UpdateElemToAll(elem)
				}
			}

		case "ID":
			connection.SetID(mrpItem.GetBody())

		case "END":
			PushElemMap()

		default:
			fmt.Println("Command Not Understood")
		}
	}
}

func handleELEMCreates(bytesMaster []byte, finalElem *elements.Element) {
	var tempElem map[string]interface{}

	json.Unmarshal(bytesMaster, &tempElem)

	test := tempElem["Components"].([]interface{})
	for _, comp := range test {

		var myComp elements.Component
		switch comp.(map[string]interface{})["Type"].(string) {

		case "SpriteRenderer":
			myComp = render.NewSpriteRenderer(finalElem, comp.(map[string]interface{})["Filename"].(string))
			finalElem.AddComponent(myComp)

		case "KeyboardMover":
			myComp = playerControl.NewKeyboardMover(finalElem, 0)
			finalElem.AddComponent(myComp)

		case "Replicator":
			myComp = playerControl.NewReplicator(finalElem, serverConnection)
			finalElem.AddComponent(myComp)

		case "Rotator":
			myComp = render.NewRotator(finalElem)
			finalElem.AddComponent(myComp)

		case "Collider":
			myComp = physics.NewCollider(finalElem)
			finalElem.AddComponent(myComp)

		default:
			fmt.Println("Component not defined")
		}
	}
	json.Unmarshal(bytesMaster, &finalElem)
}
