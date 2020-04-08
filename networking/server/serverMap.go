package server

import (
	"net"

	"github.com/jtheiss19/MARI-Engine/networking/mrp"
)

var spawingFunction func(conn net.Conn, ID string)

func sendSessionID(conn net.Conn, ID string) {
	myMRP := mrp.NewMRP([]byte("ID"), []byte(ID), []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func onSpawn(conn net.Conn, ID string) {
	spawingFunction(conn, ID)
}

func SetSpawnFunction(newSpawnFunction func(conn net.Conn, ID string)) {
	spawingFunction = newSpawnFunction
}
