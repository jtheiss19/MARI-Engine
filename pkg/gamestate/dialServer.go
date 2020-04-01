package gamestate

import (
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/mrp"
)

func Dial(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	go mrp.ReadMRPFromConn(conn, handleMRP)

	myMRP := mrp.NewMRP([]byte("test"), []byte("test"), []byte("test"))

	conn.Write(myMRP.MRPToByte())
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {

}
